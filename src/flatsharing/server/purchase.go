package main

import (
	"context"

	pb "github.com/Its-Alex/flatsharing/src/flatsharing/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListPurchase list purchases
func (s *service) ListPurchases(ctx context.Context, req *pb.ListPurchasesRequest) (*pb.ListPurchasesResponse, error) {
	return &pb.ListPurchasesResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// GetPurchase get a purchase
func (s *service) GetPurchase(ctx context.Context, req *pb.GetPurchaseRequest) (*pb.GetPurchaseResponse, error) {
	return &pb.GetPurchaseResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// CreatePurchase create a purchase
func (s *service) CreatePurchase(ctx context.Context, req *pb.CreatePurchaseRequest) (*pb.CreatePurchaseResponse, error) {
	return &pb.CreatePurchaseResponse{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// UpdatePurchase update a purchase
func (s *service) UpdatePurchase(ctx context.Context, req *pb.UpdatePurchaseRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// DeletePurchase delete a purchase
func (s *service) DeletePurchase(ctx context.Context, req *pb.DeletePurchaseRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}
