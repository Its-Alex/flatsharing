package query

import (
	"flatsharing/core/database"
	"flatsharing/core/middleware"
)

// GetUsers get all users
func GetUsers(pagination middleware.Pagination) ([]database.User, error) {
	var users []database.User
	err := database.Db.Select(&users,
		`SELECT * FROM users LIMIT $1 OFFSET $2`,
		pagination.Limit,
		pagination.Page,
	)
	return users, err
}

// GetUserByMailOrID get a user with his id or mail set
func GetUserByMailOrID(user database.User) (*database.User, error) {
	retUser := &database.User{}
	err := database.Db.Get(retUser, "SELECT * FROM users WHERE id = $1 OR mail = $2 LIMIT 1", user.ID, user.Mail)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return retUser, nil
}

// AddUser add a user to database
func AddUser(user database.User) error {
	_, err := database.Db.Exec(`INSERT INTO users(
		id,
		mail,
		login,
		username,
		password,
		firstname,
		lastname,
		role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.ID,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Role,
	)
	return err
}

// UpdateUserByID update an user with his id
func UpdateUserByID(user database.User) error {
	_, err := database.Db.Exec(`UPDATE users SET
		mail = $2,
		login = $3,
		username = $4,
		password = $5,
		firstname = $6,
		lastname = $7,
		role = $8,
		created_at = $9
		WHERE id = $1`,
		user.ID,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Role,
		user.CreatedAt,
	)
	return err
}

// DeleteUserByID delete a user with his ID
func DeleteUserByID(user database.User) error {
	_, err := database.Db.Exec("DELETE FROM users WHERE id = $1", user.ID)
	return err
}
