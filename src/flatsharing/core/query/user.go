package query

import (
	"flatsharing/core/database"
	"time"
)

func AddUser(user database.User) {
	database.Db.MustExec(`INSERT INTO users(
		id,
		mail,
		login,
		username,
		password,
		firstname,
		lastname,
		role,
		date
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		user.ID,
		user.Mail,
		user.Login,
		user.Username,
		user.Password,
		user.FirstName,
		user.LastName,
		user.Role,
		time.Now(),
	)
}

func GetUserByID(id string) database.User {
	var user database.User
	stmt, err := database.Db.Preparex(`SELECT * FROM users WHERE id = $1 LIMIT 1`)
	if err != nil {
		panic(err)
	}
	err = stmt.Get(&user, id)
	if err != nil {
		if err.Error() != "sql: no rows in result set" {
			panic(err)
		}
	}
	return user
}
