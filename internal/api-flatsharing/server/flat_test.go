package main

import (
	"testing"

	pb "github.com/Its-Alex/flatsharing/internal/api-flatsharing/v1"
	"github.com/stretchr/testify/require"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestGetFlat(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectQuery("SELECT \\* FROM flats WHERE").
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(
				"01D3HXVSBD3FF5R36GGGBZBXYB",
				"home",
			))

	res, err := s.GetFlat(ctx, &pb.GetFlatRequest{
		Flat: &pb.Flat{
			Id: "01D3HXVSBD3FF5R36GGGBZBXYB",
		},
	})

	require.Nil(t, err)
	require.Equal(t, &pb.GetFlatResponse{
		Flat: &pb.Flat{
			Id:   "01D3HXVSBD3FF5R36GGGBZBXYB",
			Name: "home",
		},
	}, res)
}

func TestCreateFlat(t *testing.T) {
	s, mock, ctx := newMockService(t)

	mock.ExpectExec("INSERT INTO flats").
		WillReturnResult(sqlmock.NewResult(1, 1))

	_, err := s.CreateFlat(ctx, &pb.CreateFlatRequest{
		Flat: &pb.Flat{
			Name: "Flat_1",
		},
	})
	require.Nil(t, err)
}
