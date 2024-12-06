package main

import (
	"context"
	"log"
	"net"
	"time"

	pb "github.com/mikej81/mini-grpc/health"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor logs each incoming request
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// Log the method being called and the request data
	log.Printf("Received request: method=%s, request=%v", info.FullMethod, req)

	// Call the handler to complete the RPC and get the response
	resp, err := handler(ctx, req)

	// Log any errors that occurred
	if err != nil {
		log.Printf("Method %s failed: %v", info.FullMethod, status.Convert(err).Message())
	} else {
		log.Printf("Method %s completed successfully", info.FullMethod)
	}

	return resp, err
}

// server implements the HealthServiceServer interface
type server struct {
	pb.UnimplementedHealthServiceServer
}

// CheckHealth returns the health status of the service
func (s *server) CheckHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	// Respect deadlines
	if deadline, ok := ctx.Deadline(); ok {
		timeLeft := time.Until(deadline)
		log.Printf("Deadline received, time left: %v", timeLeft)
	}

	// Validate the request
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	// Simulate a successful response
	log.Println("Processing CheckHealth request...")
	return &pb.HealthResponse{Status: "healthy"}, nil
}

// CheckNonHealth returns a "sick" status for demonstration purposes
func (s *server) CheckNonHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	// Respect deadlines
	if deadline, ok := ctx.Deadline(); ok {
		timeLeft := time.Until(deadline)
		log.Printf("Deadline received, time left: %v", timeLeft)
	}

	// Validate the request
	if req == nil {
		return nil, status.Errorf(codes.InvalidArgument, "request cannot be nil")
	}

	// Simulate a potential error scenario for demonstration
	if time.Now().Unix()%2 == 0 { // Simulated condition
		return nil, status.Errorf(codes.Internal, "simulated server error")
	}

	// Simulate a successful response
	log.Println("Processing CheckNonHealth request...")
	return &pb.HealthResponse{Status: "sick"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server with the logging interceptor
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(LoggingInterceptor),
	)
	pb.RegisterHealthServiceServer(grpcServer, &server{})

	// Enable gRPC Reflection
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
