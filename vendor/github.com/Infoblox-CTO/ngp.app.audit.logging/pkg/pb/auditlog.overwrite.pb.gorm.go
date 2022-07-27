package pb

import (
	"context"
)

func (m *AuditLog) AfterToORM(ctx context.Context, orm *AuditLogORM) error {
	if orm.SubjectType == SubjectType_name[0] {
		orm.SubjectType = ""
	}
	return nil
}

func (m *AuditLogRes) AfterToORM(ctx context.Context, orm *AuditLogResORM) error {
	if orm.SubjectType == SubjectType_name[0] {
		orm.SubjectType = ""
	}
	return nil
}
