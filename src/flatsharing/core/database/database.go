package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

var (
	HOST     = "localhost"
	PORT     = "5432"
	USER     = "flatsharing"
	PASSWORD = "611bukBNpbA3"
	DB_NAME  = "flatsharing"
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
