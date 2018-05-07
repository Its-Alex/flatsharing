package database

import (
	"time"
)

// User model
type User struct {
	ID        string    `json:"id"`
	Mail      string    `json:"mail"`
	Login     string    `json:"login"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Tokens    []Token   `json:"tokens,omitempty"`
	FirstName string    `json:"firstname"`
	LastName  string    `json:"lastname"`
	Role      int       `json:"role"`
	Date      time.Time `json:"date"`
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
