package main

import (
	"context"

	pb "github.com/Its-Alex/flatsharing/src/flatsharing/v1"
)

// CreatePurchase create a purchase
func (s *service) CreatePurchase(ctx context.Context, req *pb.CreatePurchaseRequest) (*pb.CreatePurchaseResponse, error) {
	return &pb.CreatePurchaseResponse{}, nil
}
