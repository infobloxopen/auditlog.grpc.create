package client

import (
	"errors"

	"time"

	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var retryCount = 10

const (
	defaultAuditLogAddr = "auditlogging:8080"
)

//AuditlogClientKey will used as key in context
type AuditlogClientKey struct{}

//type Option func(*options)
type AuditFields func(*AuditField)

// Client interface defines a Connect and close for clients to use.
type Client interface {
	Connect() error
	Close()
	Create(ctx context.Context, request *pb.CreateRequest, opts ...grpc.CallOption) (*pb.CreateResponse, error)
	FetchLogs(ctx context.Context, in *pb.DownloadRequest, opts ...grpc.CallOption) (pb.AuditLogging_FetchLogsClient, error)
}

type AuditField struct {
	pb.AuditLog
}

type AuditClient struct {
	addr   string
	conn   *grpc.ClientConn
	client *pb.AuditLoggingClient
	opts   []grpc.DialOption
}

// NewClient creates client instance using given options.
func NewClient(addr string, opts ...grpc.DialOption) Client {
	ac := &AuditClient{addr: addr, opts: opts}
	return ac
}

// Connect makes grpc connection with the server
func (c *AuditClient) Connect() (err error) {
	if c.opts == nil {
		gopts := []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithBackoffMaxDelay(3 * time.Second),
		}
		goctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		c.conn, err = grpc.DialContext(goctx, c.addr, gopts...)
	} else {
		goctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
		defer cancel()
		c.conn, err = grpc.DialContext(goctx, c.addr, c.opts...)
	}
	if err != nil {
		return err
	}

	client := pb.NewAuditLoggingClient(c.conn)
	c.client = &client
	return nil
}

func (c *AuditClient) Close() {
	if c.conn != nil {
		c.conn.Close()
		c.conn = nil
	}
	c.client = nil
}

// Create creates auditlog entry
func (c *AuditClient) Create(ctx context.Context, request *pb.CreateRequest, opts ...grpc.CallOption) (*pb.CreateResponse, error) {
	ac := c.client
	if ac == nil {
		er := errors.New("client is not initialized")
		return nil, er
	}
	resp, err := (*ac).CreateAuditLog(ctx, request, opts...)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *AuditClient) FetchLogs(ctx context.Context, in *pb.DownloadRequest, opts ...grpc.CallOption) (pb.AuditLogging_FetchLogsClient, error) {
	ac := c.client
	if ac == nil {
		er := errors.New("client is not initialized")
		return nil, er
	}
	stream, err := (*ac).FetchLogs(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return stream, nil
}

// WithAuditLoggingClient adds the client object into the context, returns a new context
func WithAuditLoggingClient(ctx context.Context, client Client) context.Context {
	return context.WithValue(ctx, AuditlogClientKey{}, client)
}

//GetAuditLogClient ...
func GetAuditLogClient(ctx context.Context) Client {
	client := ctx.Value(AuditlogClientKey{})
	return client.(*AuditClient)
}

func createDefaultAuditClient(auditlogAddr string) Client {
	count := 0
	gopts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithBackoffMaxDelay(3 * time.Second),
		grpc.WithBlock(),
	}
	for count < retryCount {
		auditClient := NewClient(auditlogAddr, gopts...)
		er := auditClient.Connect()
		if er == nil {
			return auditClient
		} else {
			count++
		}
	}
	return nil
}
