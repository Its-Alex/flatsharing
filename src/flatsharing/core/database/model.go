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
