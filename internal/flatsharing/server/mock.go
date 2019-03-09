package main

import (
	"context"
	"strings"
	"testing"

	"github.com/Its-Alex/flatsharing/internal/core/database"
	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func newMockService(t *testing.T) (*service, sqlmock.Sqlmock, context.Context) {
	s := &service{}

	// Mock database
	mockDB, mock, err := sqlmock.New()
	s.db = &database.DB{
		DB: sqlx.NewDb(mockDB, "sqlmock"),
	}
	s.db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := context.Background()

	return s, mock, ctx
}
