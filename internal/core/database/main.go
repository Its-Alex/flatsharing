package database

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"
	"github.com/spf13/viper"

	// This import is used to import postgres driver
	_ "github.com/lib/pq"
)

/*
** This file is used as a mirror on database structure
 */

// DB is the connection to database
type DB struct {
	*sqlx.DB
}

// User model
type User struct {
	ID        string `json:"id"`
	Mail      string `json:"mail" validate:"email,required"`
	Login     string `json:"login" validate:"min=5,max=32,required"`
	Username  string `json:"username" validate:"min=5,max=32,required"`
	Password  string `json:"password" validate:"min=8,max=32,required"`
	Firstname string `json:"firstname" validate:"min=5,max=32,required"`
	Lastname  string `json:"lastname" validate:"min=5,max=32,required"`
	Role      int    `json:"role"`
	CreatedAt string `json:"create_at"`
}

// Token model
type Token struct {
	ID        string `json:"id"`
	FkUserID  string `json:"fk_user_id"`
	Token     string `json:"token"`
	CreatedAt string `json:"create_at"`
}

// Flatsharing model
type Flatsharing struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Shop model
type Shop struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Purchase model
type Purchase struct {
	ID              string    `json:"id"`
	FkFlatsharingID string    `json:"flatsharing"`
	FkUserID        string    `json:"user"`
	FkBuyerID       string    `json:"buyer"`
	FkShopID        string    `json:"shop"`
	Price           int       `json:"price"`
	Description     string    `json:"description"`
	Date            time.Time `json:"date"`
}

// Connect is used to connect to a database
func Connect() (db *DB, err error) {
	maxRetry, err := strconv.Atoi(viper.GetString("psql_max_retry"))
	if err != nil {
		return nil, err
	}
	for i := 0; i < maxRetry; i++ {
		// Database setup
		var err error
		db, err := sqlx.Open("postgres", fmt.Sprintf("host=%v port=%v user=%v password=%v dbname=%v sslmode=%v",
			viper.GetString("psql_host"),
			viper.GetString("psql_port"),
			viper.GetString("psql_user"),
			viper.GetString("psql_password"),
			viper.GetString("psql_dbname"),
			viper.GetString("psql_sslmode"),
		))
		if err != nil {
			panic(err)
		}
		db.SetMaxIdleConns(10)
		if err == nil {
			db.Mapper = reflectx.NewMapperFunc("json", strings.ToLower)
			return &DB{db}, nil
		}
		log.Println("Retry to connect to PostgreSQL database")
		time.Sleep(5 * time.Second)
	}
	return nil, err
}
