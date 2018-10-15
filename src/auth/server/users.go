package main

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Its-Alex/flatsharing/src/auth/v1"
	"github.com/Its-Alex/flatsharing/src/core/database"
	"github.com/Its-Alex/flatsharing/src/core/helper"
	"github.com/Its-Alex/flatsharing/src/core/query"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	validator "gopkg.in/go-playground/validator.v9"
)

// Signin signin to server
func (s *Service) Signin(ctx context.Context, req *pb.SigninRequest) (res *pb.SigninResponse, err error) {
	return res, nil
}

// CreateUser to server
func (s *Service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (res *pb.CreateUserResponse, err error) {
	user := database.User{
		ID:   helper.GenUlid(),
		Role: 0,
		Date: time.Now(),
	}
	copier.Copy(&user, req)
	fmt.Println(user)

	if errs := validator.New().Struct(user); errs != nil {
		return nil, status.New(codes.InvalidArgument, errs.Error()).Err()
	}

	user.Password = helper.BcryptGen(user.Password)

	_, err = query.AddUser(user)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.New(codes.Internal, "Internal server error").Err()
	}

	return &pb.CreateUserResponse{
		Id: user.ID,
	}, nil
}

// GetUsers get all users
func (s *Service) GetUsers(ctx context.Context, req *empty.Empty) (res *pb.ListUsersResponse, err error) {
	return res, nil
}

// GetUser get one user
func (s *Service) GetUser(ctx context.Context, req *pb.GetUserRequest) (res *pb.GetUserResponse, err error) {
	return res, nil
}

// UpdateUser update one user
func (s *Service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (res *empty.Empty, err error) {
	return res, nil
}

// DeleteUser delete one user
func (s *Service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (res *empty.Empty, err error) {
	return res, nil
}
