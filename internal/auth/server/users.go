package main

import (
	"context"
	"database/sql"
	"strconv"

	pb "github.com/Its-Alex/flatsharing/internal/auth/v1"
	"github.com/Its-Alex/flatsharing/internal/core/helper"

	"github.com/Its-Alex/flatsharing/internal/core/database"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Signin signin to server
func (s *service) Signin(ctx context.Context, req *pb.SigninRequest) (*pb.SigninResponse, error) {
	res := &pb.SigninResponse{}

	user, err := s.db.GetUser(&pb.User{
		Mail:  req.Login,
		Login: req.Login,
	})
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if user == nil {
		helper.Logger.Error("User not found")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		if err.Error() == "crypto/bcrypt: hashedPassword is not the hash of the given password" {
			helper.Logger.Error("Password mismatch")
			return nil, status.Error(codes.InvalidArgument, "Password mismatch")
		}
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	token := &database.Token{
		ID:       helper.GenUlid(),
		FkUserID: user.Id,
		Token:    helper.GenToken(),
	}

	_, err = s.db.AddToken(token)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	res.Token = token.Token
	return res, nil
}

// ListUsers get all users
func (s *service) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	res := &pb.ListUsersResponse{}
	page, err := strconv.ParseInt(req.PageToken, 10, 32)
	if err != nil && len(req.PageToken) != 0 {
		helper.Logger.Error("Page token received is not a number: ", err)
		return nil, status.Error(codes.InvalidArgument, "page_token is not a number")
	}
	if (req.PageSize <= 0 || req.PageSize >= 200) && req.PageSize != -1 {
		req.PageSize = 50
	}

	err = s.db.Get(&res.TotalPageSize, "SELECT COUNT(*) FROM users")
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

	var rows *sqlx.Rows
	if req.PageSize == -1 {
		rows, err = s.db.Queryx("SELECT * FROM users ORDER BY id ASC")
	} else {
		rows, err = s.db.Queryx("SELECT * FROM users ORDER BY id ASC OFFSET $1 LIMIT $2", page, req.PageSize)
	}
	if err != nil {
		if err == sql.ErrNoRows {
			res.NextPageToken = ""
			return res, nil
		}
		helper.Logger.Error("Failed serching server, ", err)
		return nil, status.Error(codes.Internal, "")
	}
	for rows.Next() {
		usr := &pb.User{}
		if err = rows.StructScan(usr); err != nil {
			helper.Logger.Error("Failed scanning struct, ", err)
			return nil, status.Error(codes.Internal, "")
		}
		res.Users = append(res.Users, usr)
	}

	return res, nil
}

// GetUser get one user
func (s *service) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var err error
	res := &pb.GetUserResponse{}
	res.User, err = s.db.GetUser(req.User)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if res.User == nil {
		helper.Logger.Error("User not found")
		return nil, status.Error(codes.NotFound, "User not found")
	}
	return res, nil
}

// CreateUser to server
func (s *service) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if req.User == nil {
		helper.Logger.Error("Request with no user")
		return nil, status.Error(codes.InvalidArgument, "No user")
	}

	findUser, err := s.db.GetUser(req.User)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if findUser != nil {
		helper.Logger.Error("User already exist with same mail/login")
		return nil, status.Error(codes.InvalidArgument, "User already exist with same mail/login")
	}

	req.User.Id = helper.GenUlid()
	req.User.Password = helper.BcryptGen(req.User.Password)
	req.User.Role = 0

	_, err = s.db.AddUser(req.User)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.CreateUserResponse{
		Id: req.User.Id,
	}, nil
}

// UpdateUser update one user
func (s *service) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*empty.Empty, error) {
	res := &empty.Empty{}
	return res, nil
}

// DeleteUser delete one user
func (s *service) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*empty.Empty, error) {
	res := &empty.Empty{}
	user, err := s.db.GetUser(req.User)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}
	if user == nil {
		helper.Logger.Error("User not found")
		return nil, status.Error(codes.NotFound, "User not found")
	}

	_, err = s.db.DeleteUser(user)
	if err != nil {
		helper.Logger.Error(err)
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return res, nil
}
