package main

import (
	"context"

	"github.com/Its-Alex/flatsharing/internal/core/helper"
	pb "github.com/Its-Alex/flatsharing/internal/api-flatsharing/v1"
	"github.com/Masterminds/squirrel"
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
	var err error
	flat := new(pb.Flat)
	err = s.db.Get(flat, "SELECT * FROM flats WHERE id = $1", req.Flat.Id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			helper.Logger.Error(err)
			return nil, status.Error(codes.NotFound, "No flat with this id")
		}
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if flat == nil {
		helper.Logger.Error("Flat not found")
		return nil, status.Error(codes.NotFound, "Flat not found")
	}
	return &pb.GetFlatResponse{
		Flat: flat,
	}, nil
}

// CreateFlat create a flat
func (s *service) CreateFlat(ctx context.Context, req *pb.CreateFlatRequest) (*pb.CreateFlatResponse, error) {
	req.Flat.Id = helper.GenUlid()

	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	sql, args, err := psql.Insert("flats").
		Columns(
			"id",
			"name",
		).Values(
		req.Flat.Id,
		req.Flat.Name).ToSql()
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	_, err = s.db.Exec(sql, args...)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.CreateFlatResponse{
		Id: req.Flat.Id,
	}, nil
}

// UpdateFlat update a flat
func (s *service) UpdateFlat(ctx context.Context, req *pb.UpdateFlatRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}

// DeleteFlat delete a flat
func (s *service) DeleteFlat(ctx context.Context, req *pb.DeleteFlatRequest) (*empty.Empty, error) {
	return &empty.Empty{}, status.Error(codes.Unimplemented, "This route is not implemented yet")
}
