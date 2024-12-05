package main

import (
	"context"
	"log"
	"net"

	pb "github.com/michael/mini-grpc"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedHealthServiceServer
}

func (s *server) CheckHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Status: "healthy"}, nil
}

func (s *server) CheckNonHealth(ctx context.Context, req *pb.HealthRequest) (*pb.HealthResponse, error) {
	return &pb.HealthResponse{Status: "sick"}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHealthServiceServer(grpcServer, &server{})

	log.Println("gRPC server is running on port 50051...")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
