package auditlogging

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/infobloxopen/protoc-gen-gorm/types"
)

func ctxAuditLogFromCtx(ctx context.Context) (*ctxAuditLog, error) {
	val := ctx.Value(ctxKey)
	if cal, ok := val.(*ctxAuditLog); ok {
		return cal, nil
	}
	return nil, errors.New("Did not find auditlogging context value")
}

func Force(ctx context.Context) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Force = true
	return nil
}

func Disable(ctx context.Context) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Disable = true
	return nil
}

func SetAccountID(ctx context.Context, accountID string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.AccountId = accountID
	return nil
}

func SetAppID(ctx context.Context, appID string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.AppId = appID
	return nil
}

func SetAction(ctx context.Context, action string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.Action = action
	return nil
}

func SetMessage(ctx context.Context, message string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.Message = message
	return nil
}

func SetEventMetadata(ctx context.Context, md map[string]interface{}) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	mdJSON, err := json.Marshal(md)
	if err != nil {
		return err
	}
	cal.Audit.EventMetadata = &types.JSONValue{
		Value: string(mdJSON),
	}
	return nil
}

func SetEventVersion(ctx context.Context, version string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.EventVersion = version
	return nil
}

func SetResourceId(ctx context.Context, resourceID string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.ResourceId = resourceID
	return nil
}

func SetEventSummary(ctx context.Context, summary string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.Message = summary
	return nil
}

func SetSubjectType(ctx context.Context, subjectType int32) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.SubjectType = pb.SubjectType(subjectType)
	return nil
}

func SetSessionType(ctx context.Context, sessionType string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.SessionType = sessionType
	return nil
}

func SetSubjectGroups(ctx context.Context, subjectGroups []string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.SubjectGroups = subjectGroups
	return nil
}

func SetSessionId(ctx context.Context, sessionID string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.SessionId = sessionID
	return nil
}

func SetClientIp(ctx context.Context, clientIP string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.ClientIp = clientIP
	return nil
}

func SetUsername(ctx context.Context, userName string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.UserName = userName
	return nil
}

func SetHttpUrl(ctx context.Context, httpURL string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.HttpUrl = httpURL
	return nil
}

func SetHttpMethod(ctx context.Context, httpMethod string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.HttpMethod = httpMethod
	return nil
}

func SetHttpReqBody(ctx context.Context, httpReqBody string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.HttpReqBody = httpReqBody
	return nil
}

func SetHttpCode(ctx context.Context, httpCode int32) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.HttpCode = httpCode
	return nil
}

func SetHttpRespBody(ctx context.Context, httpRespBody string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.HttpRespBody = httpRespBody
	return nil
}

func SetCreatedAt(ctx context.Context, createdAt *timestamp.Timestamp) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.CreatedAt = createdAt
	return nil
}

func SetResult(ctx context.Context, result string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.Result = result
	return nil
}

func SetRequestId(ctx context.Context, requestID string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.RequestId = requestID
	return nil
}

func SetResourceDesc(ctx context.Context, resourceDesc string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.ResourceDesc = resourceDesc
	return nil
}

func SetResourceType(ctx context.Context, resourceType string) error {
	cal, err := ctxAuditLogFromCtx(ctx)
	if err != nil {
		return err
	}
	cal.Audit.ResourceType = resourceType
	return nil
}
