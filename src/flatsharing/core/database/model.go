package database

import (
	"time"
)

// User model
type User struct {
	ID        string    `json:"id"`
	Mail      string    `json:"mail" validate:"required,email"`
	Login     string    `json:"login" validate:"required,gte=0,lte=32"`
	Username  string    `json:"username" validate:"required,gte=0,lte=32"`
	Password  string    `json:"password" validate:"required,gte=8,lte=64"`
	Tokens    []Token   `json:"tokens,omitempty"`
	FirstName string    `json:"firstname" validate:"required,gte=0,lte=64"`
	LastName  string    `json:"lastname" validate:"required,gte=0,lte=64"`
	Role      int       `json:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// Token model
type Token struct {
	ID        string    `json:"id"`
	UserID    string    `json:"fk_user_id"`
	Token     string    `json:"token"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
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
