package main

import (
	"context"

	"github.com/Its-Alex/flatsharing/internal/core/helper"
	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/Masterminds/squirrel"
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
	req.Purchase.Id = helper.GenUlid()

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, args, err := psql.Insert("purchases").
		Columns(
			"id",
			"fk_flats_id",
			"fk_user_id",
			"fk_buyer_id",
			"fk_shop_id",
			"price",
			"description").
		Values(
			req.Purchase.Id,
			req.Purchase.FlatId,
			req.Purchase.UserId,
			req.Purchase.BuyerId,
			req.Purchase.Shop,
			req.Purchase.Price,
			req.Purchase.Desc).
		ToSql()
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.CreatePurchaseResponse{
		Id: req.Purchase.Id,
	}, nil
}

// UpdatePurchase update a purchase
func (s *service) UpdatePurchase(ctx context.Context, req *pb.UpdatePurchaseRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// DeletePurchase delete a purchase
func (s *service) DeletePurchase(ctx context.Context, req *pb.DeletePurchaseRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}
