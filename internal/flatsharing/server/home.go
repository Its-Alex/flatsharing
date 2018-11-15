package main

import (
	"context"

	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListHome list homes
func (s *service) ListHomes(ctx context.Context, req *pb.ListHomesRequest) (*pb.ListHomesResponse, error) {
	return &pb.ListHomesResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// GetHome get a home
func (s *service) GetHome(ctx context.Context, req *pb.GetHomeRequest) (*pb.GetHomeResponse, error) {
	return &pb.GetHomeResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// CreateHome create a home
func (s *service) CreateHome(ctx context.Context, req *pb.CreateHomeRequest) (*pb.CreateHomeResponse, error) {
	return &pb.CreateHomeResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// UpdateHome update a home
func (s *service) UpdateHome(ctx context.Context, req *pb.UpdateHomeRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// DeleteHome delete a home
func (s *service) DeleteHome(ctx context.Context, req *pb.DeleteHomeRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}
