package database

import (
	"flatsharing/core/helper"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Import for use database
)

// Db connection to database
var Db *sqlx.DB

var (
	HOST     = helper.GetEnv("PSQL_HOST", "localhost")
	USER     = helper.GetEnv("PSQL_USER", "flatsharing")
	PASSWORD = helper.GetEnv("PSQL_PASSWORD", "611bukBNpbA3")
	PORT     = helper.GetEnv("PSQL_PORT", "5432")
	DB_NAME  = helper.GetEnv("PSQL_DB_NAME", "flatsharing")
)

func init() {
	var err error
	Db, err = sqlx.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
		HOST,
		PORT,
		USER,
		PASSWORD,
		DB_NAME,
		"disable",
	))
	if err != nil {
		panic(err)
	}
	Db.SetMaxIdleConns(10)
}

func Ping() {
	Db.Ping()
}
