package main

import (
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	pb "github.com/Its-Alex/flatsharing/internal/api-flatsharing/v1"
	"github.com/stretchr/testify/require"
)

func TestListPurchase(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT COUNT\\(\\*\\) FROM purchases").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	mock.ExpectQuery("SELECT \\* FROM purchases ORDER BY id LIMIT 50 OFFSET 0").
		WillReturnRows(sqlmock.NewRows([]string{"id", "fk_flat_id", "fk_user_id", "fk_buyer_id", "fk_shop_id", "price", "description", "created_at"}).
			AddRow(
				"01CTR7K83ZT5XXP89S87FQHMTG",
				"02CTR7K83ZT5XXP89S87FQHMTG",
				"03CTR7K83ZT5XXP89S87FQHMTG",
				"04CTR7K83ZT5XXP89S87FQHMTG",
				"05CTR7K83ZT5XXP89S87FQHMTG",
				10,
				"Courses auchan",
				"2018-10-26 13:04:13.001629+00",
			).
			AddRow(
				"05CTR7K83ZT5XXP89S87FQHMTG",
				"02CTR7K83ZT5XXP89S87FQHMTG",
				"03CTR7K83ZT5XXP89S87FQHMTG",
				"04CTR7K83ZT5XXP89S87FQHMTG",
				"05CTR7K83ZT5XXP89S87FQHMTG",
				10,
				"Courses auchan",
				"2018-10-26 13:04:13.001629+00",
			))

	res, err := s.ListPurchases(ctx, &pb.ListPurchasesRequest{})
	require.Nil(t, err)
	require.Equal(t, int32(2), res.TotalPageSize)
	require.Equal(t, "", res.NextPageToken)
}

func TestCreatePurchase(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectExec("INSERT INTO purchases").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := s.CreatePurchase(ctx, &pb.CreatePurchaseRequest{
		Purchase: &pb.Purchase{
			FkFlatId:    "01CTR7K83ZT5XXP89S87FQHMTG",
			FkUserId:    "01CTR7K83ZT5XXP89S87FQHMTR",
			FkBuyerId:   "01CTR7K83ZT5XXP89S87FQHMTF",
			FkShopId:    "Auchan",
			Price:       11.1,
			Description: "Courses",
		},
	})
	require.Nil(t, err)
}
