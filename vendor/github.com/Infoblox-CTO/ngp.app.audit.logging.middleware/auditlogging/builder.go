package auditlogging

import (
	"fmt"
	"path"
	"reflect"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"

	"github.com/infobloxopen/atlas-app-toolkit/auth"
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	"github.com/infobloxopen/atlas-app-toolkit/requestid"
	"github.com/infobloxopen/atlas-app-toolkit/requestinfo"
	"github.com/infobloxopen/atlas-app-toolkit/rpc/resource"
	"google.golang.org/grpc/peer"
)

var (
	ErrInternal   = grpc.Errorf(codes.Internal, "Unable to process request")
	accountID     = "Default"
	GetReqDetails = getRequestDetails
)

const USER = "username"
const IDfield = "Id"
const xForwardedFor = "X-Forwarded-For"

// Builder is responsible for creating requests to Audit server.
type Builder interface {
	// build(...) uses the incoming request context to make a separate request to audit logging
	Build(ctx context.Context, req interface{}, response interface{}, respErr error) (*pb.CreateRequest, error)
}

// defaultBuilder provides a default implementation of the Builder interface
type auditlogDefaultBuilder struct{ audit pb.AuditLog }

type option func(*auditlogDefaultBuilder)

func parseError(respErr error) string {
	var message string
	parseErr, ok := status.FromError(respErr)
	if !ok {
		message = respErr.Error()
	} else {
		message = parseErr.Message()
	}
	return message
}

func parseMessage(ctx context.Context) string {
	var message string

	transportStream := grpc.ServerTransportStreamFromContext(ctx)
	v := reflect.ValueOf(transportStream).Elem().FieldByName("header")
	for _, k := range v.MapKeys() {
		if k.String() == runtime.MetadataPrefix+statusmessage {
			message = fmt.Sprintf("%v", v.MapIndex(k))
			if len(message) != 0 {
				message = message[1 : len(message)-1]
			}
			break
		}
	}
	return message
}

// NewBuilder builds a builder for making requests
func NewBuilder(opts ...option) Builder {
	db := &auditlogDefaultBuilder{
		audit: pb.AuditLog{
			Message:      "",
			Result:       "Allowed",
			AppId:        "",
			ClientIp:     "127.0.0.1",
			UserName:     "admin",
			ResourceId:   "resource-id",
			ResourceType: "resource-type",
		},
	}
	for _, opt := range opts {
		opt(db)
	}
	return db
}

func findField(reqVal reflect.Value, fieldName string) (resID resource.Identifier, found bool) {
	if reqVal.Kind() == reflect.Slice {
		if reqVal.Len() > 0 {
			reqVal = reqVal.Index(0)
		} else {
			if reqVal.Type().Elem().Kind() == reflect.Ptr {
				reqVal = reflect.New(reqVal.Type().Elem().Elem())
			} else {
				reqVal = reflect.New(reqVal.Type().Elem())
			}
		}
	}
	if reqVal.Kind() == reflect.Ptr {
		reqVal = reqVal.Elem()
	}

	for reqVal.Kind() == reflect.Ptr {
		reqVal = reqVal.Elem()
	}
	if reqVal.Kind() == reflect.Struct && !found {
		if field := reqVal.FieldByName(fieldName); field.IsValid() {
			if field.CanInterface() {
				switch t := field.Interface().(type) {
				case *resource.Identifier:
					if t == nil {
						return resID, true
					}
					return resource.Identifier{ResourceId: t.ResourceId,
						ResourceType:    t.ResourceType,
						ApplicationName: t.ApplicationName}, true
				case uint64:
					if t == 0 {
						return resource.Identifier{ResourceId: ""}, true
					}
					return resource.Identifier{ResourceId: fmt.Sprintf("%d", t)}, true
				case int64:
					if t == 0 {
						return resource.Identifier{ResourceId: ""}, true
					}
					return resource.Identifier{ResourceId: fmt.Sprintf("%d", t)}, true
				case int:
					if t == 0 {
						return resource.Identifier{ResourceId: ""}, true
					}
					return resource.Identifier{ResourceId: fmt.Sprintf("%d", t)}, true
				case string:
					return resource.Identifier{ResourceId: fmt.Sprintf("%s", t)}, true
				default:
					return resource.Identifier{ResourceId: ""}, true
				}
			} else {
				// possibly field is not exported, ignore.
			}
		} else {
			for i := 0; i < reqVal.NumField(); i++ {
				return findField(reqVal.Field(i), fieldName)
			}
		}
	}
	return resID, found
}

func findResourceFromRequest(ctx context.Context, req interface{}) (resource.Identifier, string) {
	resId, _ := findField(reflect.ValueOf(req), IDfield)
	resourceType, action, err := GetReqDetails(ctx)
	if resId.ResourceType == "" && err == nil {
		tempList := strings.Split(resourceType, ".")
		if tempList != nil {
			resType := tempList[len(tempList)-1]
			resId.ResourceType = strings.ToLower(resType)
		}
	}
	return resId, action
}

func findResourceFromResponse(resp interface{}) string {
	resId, _ := findField(reflect.ValueOf(resp), IDfield)
	return resId.GetResourceId()
}

func findFieldFromContext(ctx context.Context) (resID resource.Identifier, action string, err error) {
	requestInfo, err := requestinfo.FromContext(ctx)
	if err != nil {
		return requestInfo.Identifier, "", err
	}
	return requestInfo.Identifier, requestInfo.OperationType.String(), nil
}

func findResourceFields(ctx context.Context, req interface{}, resp interface{}) (appName, resType, resId, action string) {
	var resID resource.Identifier
	var err error

	resID, action, err = findFieldFromContext(ctx)
	if err != nil {
		resID, action = findResourceFromRequest(ctx, req)
	}
	if resID.ResourceId == "0" || resID.ResourceId == "" {
		resID.ResourceId = findResourceFromResponse(resp)
	}
	return resID.ApplicationName, resID.ResourceType, resID.ResourceId, action
}

func findIP(ctx context.Context) (string, error) {
	remoteIp, exists := gateway.Header(ctx, xForwardedFor)
	if !exists {
		p, ok := peer.FromContext(ctx)
		if !ok {
			err := status.Errorf(codes.NotFound, "Peer information not available in incoming grpc request")
			return "", err
		}
		if p != nil {
			remoteIp = strings.Split(fmt.Sprintf("%s", p.Addr), ":")[0]
		}
	} else {
		remoteIp = strings.TrimSpace(strings.Split(remoteIp, ",")[0])
	}
	return remoteIp, nil
}

// Build creating a build request
func (b *auditlogDefaultBuilder) Build(ctx context.Context, req interface{}, response interface{}, respErr error) (*pb.CreateRequest, error) {
	var appName, message string
	userName, err := auth.GetJWTField(ctx, USER, nil)
	if err != nil {
		return nil, err
	}
	b.audit.UserName = userName
	b.audit.RequestId, _ = requestid.FromContext(ctx) // RequestID is optional, hence ignoring the error
	b.audit.ClientIp, err = findIP(ctx)
	if err != nil {
		return nil, err
	}
	appName, b.audit.ResourceType, b.audit.ResourceId, b.audit.Action = findResourceFields(ctx, req, response)

	if b.audit.AppId == "" {
		if appName == "" {
			appName = "N/A"
		}
		b.audit.AppId = appName
	}
	if respErr != nil {
		message = parseError(respErr)
		b.audit.Result = Failed
	} else {
		message = parseMessage(ctx)
		b.audit.Result = Success
	}
	if b.audit.Message == "" {
		b.audit.Message = message
	}
	return &pb.CreateRequest{Payload: &b.audit}, nil
}

func getRequestDetails(ctx context.Context) (string, string, error) {
	fullMethodString, ok := grpc.Method(ctx)
	if !ok {
		return "", "", ErrInternal
	}
	return path.Dir(fullMethodString)[1:], path.Base(fullMethodString), nil
}

func WithMessage(message string) option {
	return func(o *auditlogDefaultBuilder) {
		o.audit.Message = message
	}
}

func WithApplicationId(appId string) option {
	return func(o *auditlogDefaultBuilder) {
		o.audit.AppId = appId
	}
}

func WithResult(res string) option {
	return func(o *auditlogDefaultBuilder) {
		o.audit.Result = res
	}
}

func WithResourceDescription(desc string) option {
	return func(o *auditlogDefaultBuilder) {
		o.audit.ResourceDesc = desc
	}
}
