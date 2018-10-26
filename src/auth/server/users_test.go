package main

import (
	"regexp"
	"testing"

	pb "github.com/Its-Alex/flatsharing/src/auth/v1"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestSignin(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT \\* FROM users WHERE").
		WillReturnRows(sqlmock.NewRows([]string{"id", "mail", "login", "username", "firstname", "lastname", "password", "role", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTG",
				"mail@example.com",
				"example",
				"example",
				"example_fs",
				"example_ls",
				"$2y$10$xhAPx.aBmeQxk6kq2bpWPenWA2/Ia7dYSr2ufGVky40Ip6HWZYvuW",
				0,
				"2018-10-26 13:04:13.001629+00",
			))

	mock.ExpectExec("INSERT INTO tokens").
		WillReturnResult(sqlmock.NewResult(1, 1))

	res, err := s.Signin(ctx, &pb.SigninRequest{
		Login:    "example",
		Password: "password",
	})
	require.Nil(t, err)
	require.Regexp(t, regexp.MustCompile("^[\\dA-Za-z!@#$%^&*()]{128}$"), res.Token)
}

func TestListUsers(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM users").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery("SELECT \\* FROM users ORDER BY id ASC OFFSET").
		WillReturnRows(sqlmock.NewRows([]string{"id", "mail", "login", "username", "firstname", "lastname", "password", "role", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTG",
				"mail@example.com",
				"example",
				"example",
				"example_fs",
				"example_ls",
				"$2y$10$xhAPx.aBmeQxk6kq2bpWPenWA2/Ia7dYSr2ufGVky40Ip6HWZYvuW",
				0,
				"2018-10-26 13:04:13.001629+00",
			).AddRow(
			"01CTR7K83ZT5XXP89S87FQHMTH",
			"mail2@example.com",
			"example2",
			"example2",
			"example_fs",
			"example_ls",
			"$2y$10$xhAPx.aBmeQxk6kq2bpWPenWA2/Ia7dYSr2ufGVky40Ip6HWZYvuW",
			0,
			"2018-10-26 13:04:13.001629+00",
		))

	res, err := s.ListUsers(ctx, &pb.ListUsersRequest{})
	require.Nil(t, err)
	require.Equal(t, int32(2), res.TotalPageSize)
	require.Equal(t, "", res.NextPageToken)
}

func TestGetUser(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT \\* FROM user").
		WillReturnRows(sqlmock.NewRows([]string{"id", "mail", "login", "username", "firstname", "lastname", "password", "role", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTH",
				"mail@example.com",
				"example",
				"example",
				"example_fs",
				"example_ls",
				"$2y$10$xhAPx.aBmeQxk6kq2bpWPenWA2/Ia7dYSr2ufGVky40Ip6HWZYvuW",
				0,
				"2018-10-26 13:04:13.001629+00",
			))

	res, err := s.GetUser(ctx, &pb.GetUserRequest{
		User: &pb.User{
			Id: "01CTR7K83ZT5XXP89S87FQHMTH",
		},
	})
	require.Nil(t, err)
	require.Equal(t, "01CTR7K83ZT5XXP89S87FQHMTH", res.User.Id)
}

func TestCreateUser(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT \\* FROM user").
		WillReturnRows(sqlmock.NewRows(
			[]string{"id", "mail", "login", "username", "firstname", "lastname", "password", "role", "created_at"},
		))

	mock.ExpectExec("INSERT INTO users").WillReturnResult(sqlmock.NewResult(1, 1)) // TODO improve with set of data given to insert

	res, err := s.CreateUser(ctx, &pb.CreateUserRequest{
		User: &pb.User{
			Mail:      "example@example.com",
			Login:     "example",
			Username:  "example",
			Password:  "password",
			Firstname: "example_fs",
			Lastname:  "example_ls",
		},
	})
	require.Nil(t, err)
	require.Regexp(t, regexp.MustCompile("^[\\dA-Za-z]{26}$"), res.Id)
}
