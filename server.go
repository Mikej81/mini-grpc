package main

import (
	"context"
	"fmt"
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

// CheckHealth processes the request and returns the status and a submit-string
func (s *server) CheckHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	// Respect deadlines
	if deadline, ok := ctx.Deadline(); ok {
		timeLeft := time.Until(deadline)
		log.Printf("Deadline received, time left: %v", timeLeft)
	}

	// Validate the request
	if req == nil || req.Field == "" || req.Value == "" {
		return nil, status.Errorf(codes.InvalidArgument, "field and value cannot be empty")
	}

	// Generate a submit-string based on the request
	submitString := fmt.Sprintf("Field: %s, Value: %s", req.Field, req.Value)

	// Simulate a successful response
	log.Println("Processing CheckHealth request...")
	return &pb.HealthResponse{
		Status:       "healthy",
		SubmitString: submitString,
	}, nil
}

// CheckNonHealth processes the request and returns the status and a submit-string
func (s *server) CheckNonHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	// Respect deadlines
	if deadline, ok := ctx.Deadline(); ok {
		timeLeft := time.Until(deadline)
		log.Printf("Deadline received, time left: %v", timeLeft)
	}

	// Validate the request
	if req == nil || req.Field == "" || req.Value == "" {
		return nil, status.Errorf(codes.InvalidArgument, "field and value cannot be empty")
	}

	// Generate a submit-string based on the request
	submitString := fmt.Sprintf("Field: %s, Value: %s", req.Field, req.Value)

	// Simulate a response
	log.Println("Processing CheckNonHealth request...")
	return &pb.HealthResponse{
		Status:       "sick",
		SubmitString: submitString,
	}, nil
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
