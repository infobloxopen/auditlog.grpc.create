package auditlogging

import (
	"context"
	"encoding/json"

	auditlog "github.com/Infoblox-CTO/ngp.app.audit.logging/client"
	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus/ctxlogrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	//Failed indicates method operation failed
	Failed  = "Failed"
	Success = "Success"
	//AuthKey indicates authorization key name
	AuthKey = "authorization"

	statusmessage = "status-message"

	muteKey     = "auditlog-mute-grpc"
	auditlogKey = "auditlog-values"

	grpcTrailerKey = "Grpc-Trailer-Auditlog-Values"

	logrusKey   = "middleware"
	logrusValue = "AuditLog"
)

// Logger is extracted from the context to log
var (
	Logger = ctxlogrus.Extract
)

type uniqueKey struct{}

var ctxKey = uniqueKey{}

type ctxAuditLog struct {
	Audit   *pb.AuditLog
	Force   bool
	Disable bool
}

func isTrue(v string) bool {
	return v == "true"
}

// UnaryServerInterceptor returns grpc.UnaryServerInterceptor
// that should be used as a middleware to include audit logging framework
// for logging and tracking user generated operations.
//
// It takes application id, audit logging client and list of method names to exclude
// from logging and populates audit logging with Action, Result, Application ID,
// Resource Id, Resource type, Username, Client IP, Resource Description, Message, Request ID
// and timestamp data.
func UnaryServerInterceptor(auditClient auditlog.Client, opts ...Option) grpc.UnaryServerInterceptor {
	opt := evaluateOptions(opts...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// Check if there is an http interceptor
		mute, hasHTTPKey := gateway.HeaderN(ctx, muteKey, -1)
		httpInterceptor := false
		if hasHTTPKey {
			for _, v := range mute {
				if isTrue(v) {
					httpInterceptor = true
				}
			}
		}

		// Store a log in the ctx so that application code can set values
		ctxLog := &ctxAuditLog{Audit: &pb.AuditLog{}}
		ctx = context.WithValue(ctx, ctxKey, ctxLog)

		// Run rest of interceptor chain
		resp, respErr := handler(ctx, req)

		logger := Logger(ctx)
		logger = logger.WithFields(map[string]interface{}{logrusKey: logrusValue})
		ctx = ctxlogrus.ToContext(ctx, logger)

		// Check exclusion list and "force" flag
		if !ctxLog.Force {
			if opt.excluderFunc(ctx, opt.exclusionList) || ctxLog.Disable || respErr != nil {
				// if httpInterceptor - Set response for bypassing audit logging in http layer
				if httpInterceptor {
					grpc.SetTrailer(ctx, metadata.Pairs(auditlogKey, "{\"Disable\":true}"))
				}
				return resp, respErr
			}
		}

		// Use custom Build function if available
		builder, ok := req.(Builder)
		// build a new log request
		if !ok {
			builder = NewBuilder(WithApplicationId(opt.appID))
		}

		// Extract fields for GRPC layer
		createReq, err := builder.Build(ctx, req, resp, respErr)
		if err != nil {
			logger.Warnf("unable to build the auditlog request %v", err)
			return resp, respErr
		}

		// Override fields with values specified in application layer
		ctxLog.Audit = overrideAuditFields(createReq.GetPayload(), ctxLog.Audit)

		// If HTTP interceptor - push into grpc trailer
		if httpInterceptor {
			logger.Debug("chain has http interceptor, skipping grpc")
			var dat []byte
			// passing as json seemed to save a lot of map -> object, object -> map merging code
			dat, err = json.Marshal(ctxLog)
			if err != nil {
				logger.Errorf("error marshaling log to pass to HTTP interceptor: %v", err)
				return resp, respErr
			}
			if err := grpc.SetTrailer(ctx, metadata.Pairs(auditlogKey, string(dat))); err != nil {
				logger.Error(err)
			}
			return resp, respErr
		}
		// Else - send audit log to audit logging service

		if createReq == nil {
			return resp, respErr
		}
		createReq.Payload = ctxLog.Audit
		token := metautils.ExtractIncoming(ctx).Get(AuthKey)
		auditCtx := metadata.AppendToOutgoingContext(ctx, AuthKey, token)

		// Create request
		_, err = auditClient.Create(auditCtx, createReq)
		if err != nil {
			logger.Warnf("unable to create the auditlog entry %v", err)
			return resp, respErr
		}
		logger.Debug("successfully created an audit entry")
		return resp, respErr
	}
}

// overrideAuditFields copies non-empty fields from the second arg into the first arg
// and returns the result for simplicity. This is used to override default values
// with those provided by the application or grpc layer
func overrideAuditFields(base *pb.AuditLog, with *pb.AuditLog) *pb.AuditLog {
	if with == nil {
		return base
	}
	if base == nil {
		return with
	}
	if with.AccountId != "" {
		base.AccountId = with.AccountId
	}
	if with.Action != "" {
		base.Action = with.Action
	}
	if with.AppId != "" {
		base.AppId = with.AppId
	}
	if with.ClientIp != "" {
		base.ClientIp = with.ClientIp
	}
	if with.Message != "" {
		base.Message = with.Message
	}
	if with.RequestId != "" {
		base.RequestId = with.RequestId
	}
	if with.EventMetadata != nil {
		base.EventMetadata = with.EventMetadata
	}
	if with.EventVersion != "" {
		base.EventVersion = with.EventVersion
	}
	if with.ResourceId != "" {
		base.ResourceId = with.ResourceId
	}
	if with.SubjectType != pb.SubjectType_Default {
		base.SubjectType = with.SubjectType
	}
	if with.SessionType != "" {
		base.SessionType = with.SessionType
	}
	if len(with.SubjectGroups) > 0 {
		base.SubjectGroups = with.SubjectGroups
	}
	if with.SessionId != "" {
		base.SessionId = with.SessionId
	}
	if with.UserName != "" {
		base.UserName = with.UserName
	}
	if with.HttpUrl != "" {
		base.HttpUrl = with.HttpUrl
	}
	if with.HttpMethod != "" {
		base.HttpMethod = with.HttpMethod
	}
	if with.HttpReqBody != "" {
		base.HttpReqBody = with.HttpReqBody
	}
	if with.HttpCode != 0 {
		base.HttpCode = with.HttpCode
	}
	if with.HttpRespBody != "" {
		base.HttpRespBody = with.HttpRespBody
	}
	if with.CreatedAt != nil {
		base.CreatedAt = with.CreatedAt
	}
	if with.Result != "" {
		base.Result = with.Result
	}
	if with.RequestId != "" {
		base.RequestId = with.RequestId
	}
	if with.ResourceDesc != "" {
		base.ResourceDesc = with.ResourceDesc
	}
	if with.ResourceType != "" {
		base.ResourceType = with.ResourceType
	}
	return base
}
