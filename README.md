# auditlog.grpc.create

The purpose of this repo is to provide a simple client for sending an auditlog GRPC Create request.

To send a request:

1) port-forward to an auditlog pod on the environment that you want to send a request on, like so:
```kb port-forward -n auditlog <auditlog-pod-name> 9090:9090```

2) edit the `jwt` variable on line 18 of `main.go` to include a valid jwt for the account you want to send a request on

3) (Optional) edit the `req` variable on lines 20-24 to modify the payload of the request.

4) `go run cmd/server/main.go`