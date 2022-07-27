package auditlog

import (
	"fmt"
	"time"

	"google.golang.org/grpc"

	log "github.com/sirupsen/logrus"

	"github.com/Infoblox-CTO/ngp.app.audit.logging/client"
)

const (
	RetryTimeoutErr = `Number of attempts to connect to audit logging service exceeded the specified limit: %v`
)

// CreateAuditClient function initializes audit logging service client and tries
// to connect to it <retry> times using blocking connection with 3s wait interval.
func NewClient(addr string, retry int) (client.Client, error) {
	log.Debugf("auditlogging: connecting to audit logging service %s...", addr)
	count := 0

	gopts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(3 * time.Second),
		grpc.WithBlock(),
	}

	var err error
	for count < retry {
		auditClient := client.NewClient(addr, gopts...)

		err = auditClient.Connect()
		if err == nil {
			return auditClient, nil
		} else {
			count++
			log.Debugf("[AuditLogging] Audit logging service connection attempt %d failed.", count)
		}
	}

	return nil, fmt.Errorf(RetryTimeoutErr, err)
}
