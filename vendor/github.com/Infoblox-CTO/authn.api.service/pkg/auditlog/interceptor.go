package auditlog

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	log "github.com/sirupsen/logrus"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/grpc-ecosystem/go-grpc-middleware/util/metautils"

	"github.com/Infoblox-CTO/ngp.app.audit.logging.middleware/auditlogging"
	"github.com/Infoblox-CTO/ngp.app.audit.logging/client"
	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
)

var (
	AuthKey         = "authorization"
	HTTPMethodMDKey = `auditlog-http-method`
)

func MetadataAnnotator(ctx context.Context, r *http.Request) metadata.MD {
	md := make(metadata.MD)
	md.Set(HTTPMethodMDKey, r.Method)
	return md
}

func UnaryServerInterceptor(c client.Client, app string, exclusion auditlogging.ExclusionList) grpc.UnaryServerInterceptor {

	var defaultBuilder = auditlogging.NewBuilder(auditlogging.WithApplicationId(app))

	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, respErr error) {

		var (
			l   *Logger
			err error
		)

		l = NewLogger()

		if resp, respErr = handler(WithLogger(ctx, l), req); respErr != nil || ExcludeFunc(ctx, exclusion) {
			return resp, respErr
		}

		builder, ok := req.(auditlogging.Builder)
		if !ok {
			builder = defaultBuilder
		}

		createReq, err := builder.Build(ctx, req, resp, respErr)
		if err != nil {
			log.Errorf(`auditlogging: builder.Build: %v`, err)
			return resp, respErr
		}

		token := metautils.ExtractIncoming(ctx).Get(AuthKey)
		if token == "" {
			log.Errorf(`auditlogging: unable to find AuthKey`)
			return resp, respErr
		}

		auditCtx := metadata.AppendToOutgoingContext(ctx, AuthKey, token)

		action := ActionFromHTTP(metautils.ExtractIncoming(ctx).Get(HTTPMethodMDKey))

		for _, r := range l.Resources {

			var payload = *createReq.Payload

			payload.ResourceId = r.GetAuditID()
			payload.ResourceType = r.GetAuditType()
			payload.Action = action.String()
			payload.Result = auditlogging.Success

			if r.GetAuditName() != "" {
				payload.Message = fmt.Sprintf(action.MessageFormat(), r.GetAuditDisplayType(), r.GetAuditName())
			} else {
				payload.Message = fmt.Sprintf("%s has been %s", r.GetAuditDisplayType(), strings.ToLower(action.Status()))
			}

			if _, err = c.Create(auditCtx, &pb.CreateRequest{Payload: &payload}); err != nil {
				log.Warnf(`auditlogging: unable to create the auditlog entry: %v`, err)
				return resp, respErr
			}
		}

		return resp, respErr
	}
}

func ExcludeFunc(ctx context.Context, exclusion auditlogging.ExclusionList) bool {

	if len(exclusion) == 0 {
		return false
	}

	service, method, err := auditlogging.GetReqDetails(ctx)
	if err != nil {
		log.Errorf(`auditlogging: couldn't get the request details: %v`, err)
		return true
	}

	method = service + "." + method
	for _, op := range exclusion {
		if op != "" && (op == service || op == method) {
			return true
		}
	}

	return false
}
