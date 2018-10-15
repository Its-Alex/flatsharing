package database

import (
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Db is the connection to database
var Db *sqlx.DB

// User model
type User struct {
	ID        string    `json:"id"`
	Mail      string    `json:"mail" validate:"email,required"`
	Login     string    `json:"login" validate:"min=5,max=32,required"`
	Username  string    `json:"username" validate:"min=5,max=32,required"`
	Password  string    `json:"password" validate:"min=8,max=32,required"`
	Firstname string    `json:"firstname" validate:"min=5,max=32,required"`
	Lastname  string    `json:"lastname" validate:"min=5,max=32,required"`
	Role      int       `json:"role"`
	Date      time.Time `json:"date"`
	Tokens    []Token   `json:"tokens,omitempty"`
}

// Token model
type Token struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	Date  string `json:"date"`
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
