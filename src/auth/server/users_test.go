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
