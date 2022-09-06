# grpc-health-check

gRPC health check minimum server

### Run

```sh
$ go run cmd/server/main.go
```

### client

```sh
$ grpcurl -plaintext -d '{"service": "health-check-test"}' localhost:50051 grpc.health.v1.Health.Check
{
  "status": "SERVING"
}
```

