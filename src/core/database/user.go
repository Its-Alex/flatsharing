package database

import (
	"database/sql"

	pb "github.com/Its-Alex/flatsharing/src/auth/v1"
	"github.com/Its-Alex/flatsharing/src/core/middleware"
)

// GetUsers get all users
func (db *DB) GetUsers(pagination middleware.Pagination) ([]User, error) {
	var users []User
	err := db.Get(&users,
		`SELECT * FROM users LIMIT $1 OFFSET $2`,
		pagination.Limit,
		pagination.Page,
	)
	return users, err
}

// AddUser add a user to database
func (db *DB) AddUser(user *pb.User) (sql.Result, error) {
	return db.Exec(`INSERT INTO users(
		id,
		mail,
		login,
		username,
		password,
		firstname,
		lastname,
		role
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		user.Id,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.Firstname,
		user.Lastname,
		user.Role,
	)
}

// GetUser get a user with his id
func (db *DB) GetUser(userData *pb.User) (*pb.User, error) {
	user := new(pb.User)
	err := db.Get(user,
		`SELECT * FROM users WHERE id = $1 OR login = $2 OR mail = $3`,
		userData.Id,
		userData.Login,
		userData.Mail,
	)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}
	return user, err
}

// UpdateUserByID update an user with his id
func (db *DB) UpdateUserByID(user User) (sql.Result, error) {
	return db.Exec(`UPDATE users SET
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

// DeleteUser update an user with his id
func (db *DB) DeleteUser(user *pb.User) (sql.Result, error) {
	return db.Exec(`DELETE FROM users
		WHERE id = $1 OR login = $2 OR mail = $3`,
		user.Id,
		user.Login,
		user.Mail,
	)
}
