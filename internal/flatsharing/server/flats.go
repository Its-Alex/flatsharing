package main

import (
	"context"

	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListFlat list flats
func (s *service) ListFlats(ctx context.Context, req *pb.ListFlatsRequest) (*pb.ListFlatsResponse, error) {
	return &pb.ListFlatsResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// GetFlat get a flat
func (s *service) GetFlat(ctx context.Context, req *pb.GetFlatRequest) (*pb.GetFlatResponse, error) {
	return &pb.GetFlatResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// CreateFlat create a flat
func (s *service) CreateFlat(ctx context.Context, req *pb.CreateFlatRequest) (*pb.CreateFlatResponse, error) {
	return &pb.CreateFlatResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// UpdateFlat update a flat
func (s *service) UpdateFlat(ctx context.Context, req *pb.UpdateFlatRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// DeleteFlat delete a flat
func (s *service) DeleteFlat(ctx context.Context, req *pb.DeleteFlatRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}
