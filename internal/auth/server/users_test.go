package main

import (
	"regexp"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Its-Alex/flatsharing/internal/auth/v1"
	"github.com/stretchr/testify/require"
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

func TestLogout(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT \\* FROM tokens WHERE").
		WithArgs("asdasdasdajwljdlj1i2ji31u8dua8djasd9aiwd9").
		WillReturnRows(sqlmock.NewRows([]string{"id", "fk_user_id", "token", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTG",
				"01CTR7K83ZT5XXP89S87FQHMTT",
				"asdasdasdajwljdlj1i2ji31u8dua8djasd9aiwd9",
				"2018-10-26 13:04:13.001629+00",
			))

	mock.ExpectQuery("SELECT \\* FROM users WHERE").
		WillReturnRows(sqlmock.NewRows([]string{"id", "mail", "login", "username", "firstname", "lastname", "password", "role", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTT",
				"mail@example.com",
				"example",
				"example",
				"example_fs",
				"example_ls",
				"$2y$10$xhAPx.aBmeQxk6kq2bpWPenWA2/Ia7dYSr2ufGVky40Ip6HWZYvuW",
				0,
				"2018-10-26 13:04:13.001629+00",
			))

	mock.ExpectExec("DELETE FROM tokens").
		WithArgs("asdasdasdajwljdlj1i2ji31u8dua8djasd9aiwd9").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := s.Logout(ctx, &pb.LogoutRequest{
		Token: "asdasdasdajwljdlj1i2ji31u8dua8djasd9aiwd9",
	})
	require.Nil(t, err)
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

func TestDeleteUser(t *testing.T) {
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

	mock.ExpectExec("DELETE FROM users").WillReturnResult(sqlmock.NewResult(1, 1)) // TODO improve with set of data given to insert

	_, err := s.DeleteUser(ctx, &pb.DeleteUserRequest{
		User: &pb.User{
			Id: "01CTR7K83ZT5XXP89S87FQHMTH",
		},
	})
	require.Nil(t, err)
}
