package pb

import (
	"github.com/infobloxopen/atlas-app-toolkit/gateway"
	_ "github.com/infobloxopen/atlas-app-toolkit/rpc/errdetails"
)

func init() {
	forward_AuditLogging_ListAuditLogs_0 = gateway.ForwardResponseMessage
}
