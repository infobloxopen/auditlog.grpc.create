package auditlogging

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/infobloxopen/atlas-app-toolkit/requestid"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"

	auditlog "github.com/Infoblox-CTO/ngp.app.audit.logging/client"
	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
)

const (
	defaultHttpBodyLimit = 3000
	xOriginalURI         = "X-Original-URI"
	trailerHeader        = "Trailer"
)

func NewHTTPMiddleware(client auditlog.Client, opts ...HTTPMiddlewareOption) func(http.Handler) http.Handler {
	mdw := newHTTPMdw(client)
	for _, opt := range opts {
		opt(mdw)
	}
	return mdw.httpMiddleware
}

type HTTPMiddlewareOption func(*httpMdw)

func newHTTPMdw(client auditlog.Client) *httpMdw {
	return &httpMdw{
		client:        client,
		logger:        logrus.NewEntry(logrus.StandardLogger()),
		httpBodyLimit: defaultHttpBodyLimit,
	}
}

type httpMdw struct {
	client        auditlog.Client
	logger        *logrus.Entry
	hook          func(*http.Request, *pb.CreateRequest) error
	httpBodyLimit int
}

func (mdw *httpMdw) httpMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		// Store map in ctx so that application code can set values
		ctxLog := &ctxAuditLog{Audit: &pb.AuditLog{}}
		req = req.WithContext(context.WithValue(req.Context(), ctxKey, ctxLog))
		reqId := req.Header.Get(requestid.DefaultRequestIDKey)
		logger := mdw.logger.WithFields(map[string]interface{}{requestid.DefaultRequestIDKey: reqId, logrusKey: logrusValue})

		// Set a grpc header suppressing logging at the grpc level in case this in a grpc-gateway
		req.Header.Add("Grpc-Metadata-"+muteKey, "true")

		// a customResponseWriter overrides the default functions to save a few extra
		// values for use later and drop the value passed from GRPC (if present) before
		// it's sent in the response to the user
		crw := &customResponseWriter{
			ResponseWriter: resp,
			httpBodyLimit:  mdw.httpBodyLimit,
		}
		if flusher, ok := resp.(http.Flusher); ok {
			crw.Flusher = flusher
		}
		reqBody, err := ioutil.ReadAll(req.Body)
		if err != nil {
			logger.Errorf("error reading reqBody: %v", err)
		}
		req.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		handler.ServeHTTP(crw, req)

		// Actual trailer values are available now, if the grpc auditlog trailer is
		// written, it overrides the http values, as it is expected that no custom
		// http is executed in that case, only grpc code
		if vals, ok := resp.Header()[grpcTrailerKey]; ok && len(vals) > 0 {
			err := json.Unmarshal([]byte(vals[0]), ctxLog)
			if err != nil {
				logger.Errorf("json unmarshal error parsing from GRPC interceptor %v", err)
				return
			}
			logger.Debugf("received partial auditlog from GRPC interceptor %+v", ctxLog)
		}

		// check service value to see if logging forced
		if !ctxLog.Force {
			// Check for muting
			if ctxLog.Disable {
				logger.Debugf("disabled, skipping HTTP interceptor")
				return
			}
			// Suppress for failed calls
			if crw.code/100 != 2 {
				logger.Debug("skipping auditing for failed call")
				return
			}
			switch req.Method {
			case "POST":
			case "PUT":
			case "PATCH":
			case "DELETE":
			default:
				logger.Debug("skipping auditing for non POST/PUT/PATCH/DELETE request")
				return
			}
		}

		// Create and send AuditLog entry
		entry := &pb.CreateRequest{
			Payload: &pb.AuditLog{},
		}

		populateDefaultHTTPData(entry.Payload, crw, req)
		if len(reqBody) > mdw.httpBodyLimit {
			entry.Payload.HttpReqBody = string(reqBody[:mdw.httpBodyLimit])
		} else {
			entry.Payload.HttpReqBody = string(reqBody)
		}
		entry.Payload.RequestId = reqId

		// Override with values from application or grpc layer
		entry.Payload = overrideAuditFields(entry.Payload, ctxLog.Audit)

		// Run hook to modify HTTP request and response presented for auditlogging
		if mdw.hook != nil {
			err := mdw.hook(req, entry)
			if err != nil {
				logger.Warnf("got error from hook: %v", err)
				return
			}
		}

		// send to audit logging service
		token := req.Header.Get(AuthKey)
		auditCtx := metadata.AppendToOutgoingContext(context.Background(), AuthKey, token)
		_, err = mdw.client.Create(auditCtx, entry)
		if err != nil {
			logger.Warnf("unable to create the auditlog entry %v", err)
			return
		}
		logger.Debug("successfully created an audit entry")
	})
}

type customResponseWriter struct {
	http.ResponseWriter
	http.Flusher
	trailerKeys   []string
	data          []byte
	code          int
	httpBodyLimit int
}

// WriteHeader removes the audit logging key from the trailer before writing it
// so that it won't be passed as part of the http response, as well as saving
// the status code of the call
func (crw *customResponseWriter) WriteHeader(statusCode int) {
	header := crw.ResponseWriter.Header()
	trailer := header[trailerHeader]
	// Remove auditlog key from trailer so it won't be returned in HTTP response
	for i, t := range trailer {
		if t == grpcTrailerKey {
			trailer = append(trailer[:i], trailer[i+1:]...)
			break
		}
	}
	crw.ResponseWriter.Header()[trailerHeader] = trailer
	crw.code = statusCode
	crw.ResponseWriter.WriteHeader(statusCode)
}

// Write saves a copy of the response data for the audit log building process
func (crw *customResponseWriter) Write(data []byte) (int, error) {
	if crw.code == 0 {
		crw.WriteHeader(http.StatusOK)
	}
	space := crw.httpBodyLimit - len(crw.data)
	if len(data) > space {
		crw.data = append(crw.data, data[:space]...)
	} else {
		crw.data = append(crw.data, data...)
	}
	return crw.ResponseWriter.Write(data)
}

func populateDefaultHTTPData(al *pb.AuditLog, crw *customResponseWriter, req *http.Request) {
	al.HttpMethod = req.Method
	al.HttpCode = int32(crw.code)
	al.HttpRespBody = string(crw.data)
	if h := req.Header.Get(xOriginalURI); h != "" {
		al.HttpUrl = h
	} else {
		al.HttpUrl = req.URL.String()
	}
	if h := req.Header.Get(xForwardedFor); h != "" {
		al.ClientIp = strings.TrimSpace(strings.Split(h, ",")[0])
	}
	if crw.code/100 == 2 {
		al.Result = Success
	} else {
		al.Result = Failed
	}
}

func WithLogger(logger *logrus.Entry) HTTPMiddlewareOption {
	return func(mdw *httpMdw) {
		mdw.logger = logger
	}
}

// WithHook capable of modifying HTTP request and response presented for auditing
func WithHook(hook func(*http.Request, *pb.CreateRequest) error) HTTPMiddlewareOption {
	return func(mdw *httpMdw) {
		mdw.hook = hook
	}
}

// WithHTTPBodyLimit sets the HTTP body limit
func WithHTTPBodyLimit(bodyLimit int) HTTPMiddlewareOption {
	return func(mdw *httpMdw) {
		mdw.httpBodyLimit = bodyLimit
	}
}
