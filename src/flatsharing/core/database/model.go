package database

type User struct {
	ID        string  `json:"id" pg:"bit"`
	Mail      string  `json:"mail" pg:"varchar(36)"`
	FirstName string  `json:"firstname" pg:"varchar(36)"`
	LastName  string  `json:"lastname" pg:"varchar(36)"`
	Username  string  `json:"username" pg:"varchar(36)"`
	Password  string  `json:"password" pg:"varchar(36)"`
	Tokens    []Token `json:"tokens,omitempty"`
	Role      int     `json:"role" pg:"int"`
	Date      int64   `json:"date" pg:"timestamptz"`
}

type Token struct {
	ID    string `json:"id" pg:"bit"`
	Token string `json:"token" pg:"varchar(128)"`
	Date  string `json:"date" pg:"timestamptz"`
}
