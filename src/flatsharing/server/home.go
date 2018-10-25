package main

import (
	"context"

	pb "github.com/Its-Alex/flatsharing/src/flatsharing/v1"
)

// CreateHome create a home
func (s *service) CreateHome(ctx context.Context, req *pb.CreateHomeRequest) (*pb.CreateHomeResponse, error) {
	return &pb.CreateHomeResponse{}, nil
}
