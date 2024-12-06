# mini-grpc

Just a mini gRPC tester

## Mini-gRPC Server

This project demonstrates a simple gRPC server with two endpoints: /health and /non-health. The server is written in Go and exposes the following functionality:

- CheckHealth: Responds with healthy.
- CheckNonHealth: Responds with sick.

## Getting Started

### Prerequisites

Go: Install Go (1.19+ recommended).

protoc: Install the Protocol Buffers compiler and the Go plugins:

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

grpcurl: Install grpcurl for testing:

```bash
git clone https://github.com/mikej81/mini-grpc.git
cd mini-grpc
```

Generate the gRPC code:

```bash
protoc --go_out=. --go-grpc_out=. health.proto
```

Start the gRPC server:

```bash
go run server.go
```

The server will run on localhost:50051.

## Testing the Server

You can test the gRPC server using grpcurl commands.

### CheckHealth Endpoint

Invoke the CheckHealth method:

```bash
grpcurl -plaintext localhost:50051 health.HealthService/CheckHealth
```

Expected output:

json
Copy code
{
  "status": "healthy"
}

### CheckNonHealth Endpoint

Invoke the CheckNonHealth method:

```bash
grpcurl -plaintext localhost:50051 health.HealthService/CheckNonHealth
```

Expected output:

```json
{
  "status": "sick"
}
```

List Available Services

List all services exposed by the gRPC server:

```bash
grpcurl -plaintext localhost:50051 list
```

Expected output:

```bash
health.HealthService
grpc.reflection.v1alpha.ServerReflection
```

List Methods of a Service

List methods provided by the health.HealthService:

```bash
grpcurl -plaintext localhost:50051 list health.HealthService
```

Expected output:

```bash
health.HealthService.CheckHealth
health.HealthService.CheckNonHealth
```

### Additional Notes

The server uses gRPC Reflection for dynamic discovery of services and methods.
For any issues, ensure the server is running and accessible at localhost:50051.
