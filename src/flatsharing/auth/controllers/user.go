package controllers

import (
	"flatsharing/core/database"
	"flatsharing/core/helper"
	"flatsharing/core/middleware"
	"flatsharing/core/query"
	"fmt"
	"time"

	"github.com/labstack/echo"
)

// GetUsers get all users
func GetUsers(c echo.Context) error {
	users := query.GetUsers(c.Get("pagination").(middleware.Pagination))
	if len(users) == 0 {
		c.Set("return_code", 204)
	} else {
		c.Set("return_code", 200)
	}
	c.Set("return_error", helper.Error{})
	c.Set("return", users)
	return nil
}

// GetUser get an user
func GetUser(c echo.Context) error {
	return nil
}

// AddUser add an user
func AddUser(c echo.Context) error {
	user := database.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return err
	}

	user.ID = helper.GenUlid()
	user.Date = time.Now()

	query.AddUser(user)
	c.Set("return_code", 200)
	c.Set("return_error", helper.Error{})
	c.Set("return", user)
	return nil
}

// UpdateUser update an user
func UpdateUser(c echo.Context) error {
	return nil
}
