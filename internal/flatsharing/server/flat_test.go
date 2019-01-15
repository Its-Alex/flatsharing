package main

import (
	"testing"

	pb "github.com/Its-Alex/flatsharing/internal/flatsharing/v1"
	"github.com/stretchr/testify/require"
	sqlmock "gopkg.in/DATA-DOG/go-sqlmock.v1"
)

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
