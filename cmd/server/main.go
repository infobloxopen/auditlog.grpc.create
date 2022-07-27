package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"net"

	auditlog "github.com/Infoblox-CTO/ngp.app.audit.logging/client"
	"github.com/Infoblox-CTO/ngp.app.audit.logging/pkg/pb"

	"github.com/sirupsen/logrus"
)

// Include a valid jwt for the account you want to create an audit log record for
var jwt = "Bearer <INSERT JWT HERE>"

var req = &pb.CreateRequest{Payload: &pb.AuditLog{ResourceDesc: "", Message: "", ClientIp: "127.0.0.1",
		ResourceId: "", RequestId: "123456789",
		HttpReqBody: `{"Username":"Mr. Magoo", "Password":"hamburger", "Id":"1337", "apikey":"abc1234", "client_ip":"127.0.0.1"}`,
		HttpRespBody: `{"Username":"Mr. Magoo", "Password":"hamburger", "Id":"1337", "apikey":"abc1234", "client_ip":"127.0.0.1"}`,
		UserName: "cheese@burger.today", Action: "Create", AppId: "Default", ResourceType: "servicetype", Result: "Success"}}

type FakeStream struct {
	header map[string]string
}

func (FakeStream) Method() string                  { return "/service.NotificationsDelivery/List" }
func (FakeStream) SetHeader(md metadata.MD) error  { return nil }
func (FakeStream) SendHeader(md metadata.MD) error { return nil }
func (FakeStream) SetTrailer(md metadata.MD) error { return nil }

func main() {
	// kb port-forward -n auditlog <auditlog-pod-name> 9090:9090
	// port-forward to audit log pod and then connect to localhost
	client := auditlog.NewClient("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	err := client.Connect()
	if err != nil {
		logrus.Errorf("Had an error connecting %s", err.Error())
	}

	logrus.Infof("Producing auditlog: %+v", req)

	if _, err = client.Create(initializeContext(), &pb.CreateRequest{Payload: req.Payload}); err != nil {
		logrus.Errorf("Had an error %s", err.Error())
	}
}

func initializeContext() context.Context {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", jwt)
	ctx = metadata.NewIncomingContext(ctx, metadata.MD{"authorization": []string{jwt}})
	ctx = peer.NewContext(ctx, &peer.Peer{
		Addr: &net.IPAddr{IP: net.ParseIP("1.2.3.4"), Zone: "Foo"},
	})
	ctx = grpc.NewContextWithServerTransportStream(ctx, &FakeStream{header: map[string]string{"authorization": "foo"}})

	return ctx
}