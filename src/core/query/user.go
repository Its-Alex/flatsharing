package query

import (
	"database/sql"

	"github.com/Its-Alex/flatsharing/src/core/database"
	"github.com/Its-Alex/flatsharing/src/core/middleware"
)

// GetUsers get all users
func GetUsers(pagination middleware.Pagination) ([]database.User, error) {
	var users []database.User
	err := database.Db.Get(&users,
		`SELECT * FROM users LIMIT $1 OFFSET $2`,
		pagination.Limit,
		pagination.Page,
	)
	return users, err
}

// AddUser add a user to database
func AddUser(user database.User) (sql.Result, error) {
	return database.Db.Exec(`INSERT INTO users(
		id,
		mail,
		login,
		username,
		password,
		firstname,
		lastname,
		role,
		created_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		user.ID,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Role,
		user.Date,
	)
}

// GetUserByID get a user with his id
func GetUserByID(id string) (database.User, error) {
	var user database.User
	err := database.Db.Get(&user,
		`SELECT * FROM users WHERE id = ?`,
		id,
	)
	return user, err
}

// UpdateUserByID update an user with his id
func UpdateUserByID(user database.User) (sql.Result, error) {
	return database.Db.Exec(`UPDATE users SET
		mail = $2,
		login = $3,
		username = $4,
		password = $5,
		firstname = $6,
		lastname = $7,
		role = $8,
		date = $9
		WHERE id = $1`,
		user.ID,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Role,
		user.Date,
	)
}
