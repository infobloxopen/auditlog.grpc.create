package auditlog

import (
	"context"
)

type AuditLoggerCtxKey struct{}

// Resource ...
type Resource interface {

	// GetID ...
	GetAuditID() string

	// GetAuditName ...
	GetAuditName() string

	// GetAuditType ...
	GetAuditType() string

	// GetAuditDisplayType ...
	GetAuditDisplayType() string
}

type Logger struct {
	Resources []Resource
}

func NewLogger() *Logger {
	return &Logger{}
}

func WithLogger(ctx context.Context, l *Logger) context.Context {
	return context.WithValue(ctx, AuditLoggerCtxKey{}, l)
}

func GetLoggerFromContext(ctx context.Context) *Logger {
	l, _ := ctx.Value(AuditLoggerCtxKey{}).(*Logger)
	return l
}

func WithResource(ctx context.Context, r ...Resource) {
	l := GetLoggerFromContext(ctx)
	if l == nil {
		return
	}

	l.Resources = append(l.Resources, r...)
	return
}
