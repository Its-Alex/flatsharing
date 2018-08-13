package query

import (
	"flatsharing/core/database"
	"flatsharing/core/helper"
	"time"
)

// GetTokenByID get a user with his id
func GetTokenByID(token database.Token) (*database.Token, error) {
	retToken := &database.Token{}
	err := database.Db.Get(retToken, "SELECT * FROM tokens WHERE id = $1 LIMIT 1", token.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return retToken, nil
}

// CreateToken create a token for specified user
func CreateToken(user database.User) (database.Token, error) {
	token := database.Token{
		ID:        helper.GenUlid(),
		UserID:    user.ID,
		Token:     helper.GenToken(),
		CreatedAt: time.Now(),
	}
	_, err := database.Db.Exec(`INSERT INTO tokens(
		id,
		fk_user_id,
		token,
		created_at
		) VALUES ($1, $2, $3, $4)`,
		token.ID,
		token.UserID,
		token.Token,
		token.CreatedAt,
	)
	return token, err
}

// GetTokenByUserID get a user with his id
func GetTokenByUserID(user database.Token) (*database.Token, error) {
	retToken := &database.Token{}
	err := database.Db.Get(retToken, "SELECT * FROM tokens WHERE fk_user_id = $1 LIMIT 1", user.ID)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return retToken, nil
}
