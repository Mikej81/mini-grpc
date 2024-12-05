package main

import (
	"context"
	"testing"

	pb "github.com/mikej81/mini-grpc/health"
)

func TestCheckHealth(t *testing.T) {
	s := &server{}
	resp, err := s.CheckHealth(context.Background(), &pb.HealthRequest{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Status != "healthy" {
		t.Fatalf("Expected 'healthy', got %s", resp.Status)
	}
}

func TestCheckNonHealth(t *testing.T) {
	s := &server{}
	resp, err := s.CheckNonHealth(context.Background(), &pb.HealthRequest{})
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if resp.Status != "sick" {
		t.Fatalf("Expected 'sick', got %s", resp.Status)
	}
}
