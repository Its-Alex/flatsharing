package main

import (
	"testing"

	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestCreatePurchase(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectExec("INSERT INTO purchases").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := s.CreatePurchase(ctx, &pb.CreatePurchaseRequest{
		Purchase: &pb.Purchase{
			FlatId:  "01CTR7K83ZT5XXP89S87FQHMTG",
			UserId:  "01CTR7K83ZT5XXP89S87FQHMTR",
			BuyerId: "01CTR7K83ZT5XXP89S87FQHMTF",
			Shop:    "Auchan",
			Price:   11.1,
			Desc:    "Courses",
		},
	})
	require.Nil(t, err)
}
