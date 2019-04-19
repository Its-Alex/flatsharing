package main

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/Its-Alex/flatsharing/internal/core/helper"
	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/Masterminds/squirrel"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListPurchase list purchases
func (s *service) ListPurchases(ctx context.Context, req *pb.ListPurchasesRequest) (*pb.ListPurchasesResponse, error) {
	res := &pb.ListPurchasesResponse{}
	page, err := strconv.ParseInt(req.PageToken, 10, 32)
	if err != nil && len(req.PageToken) != 0 {
		helper.Logger.Error("Page token received is not a number: ", err)
		return nil, status.Error(codes.InvalidArgument, "page_token is not a number")
	}
	if (req.PageSize <= 0 || req.PageSize >= 200) && req.PageSize != -1 {
		req.PageSize = 50
	}

	err = s.db.Get(&res.TotalPageSize, "SELECT COUNT(*) FROM purchases")
	if err != nil {
		helper.Logger.Error("Failed couting server, ", err)
		return nil, status.Error(codes.Internal, "")
	}
	if res.TotalPageSize <= int32(page) {
		return res, nil
	}
	if res.TotalPageSize >= int32(page)+req.PageSize {
		res.NextPageToken = strconv.FormatInt(page+int64(req.PageSize), 10)
	}

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	query, _, err := psql.
		Select("*").
		From("purchases").
		OrderBy("id").
		Offset(uint64(page)).
		Limit(uint64(req.PageSize)).
		ToSql()
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	var rows *sqlx.Rows
	rows, err = s.db.Queryx(query)
	if err != nil {
		if err == sql.ErrNoRows {
			res.NextPageToken = ""
			return res, nil
		}
		helper.Logger.Error("Failed serching server, ", err)
		return nil, status.Error(codes.Internal, "")
	}
	for rows.Next() {
		usr := &pb.Purchase{}
		if err = rows.StructScan(usr); err != nil {
			helper.Logger.Error("Failed scanning struct, ", err)
			return nil, status.Error(codes.Internal, "")
		}
		res.Purchase = append(res.Purchase, usr)
	}

	return res, nil
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
			req.Purchase.FkFlatId,
			req.Purchase.FkUserId,
			req.Purchase.FkBuyerId,
			req.Purchase.FkShopId,
			req.Purchase.Price,
			req.Purchase.Description).
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
