package controllers

import (
	"flatsharing/core/database"
	"flatsharing/core/helper"
	"flatsharing/core/middleware"
	"flatsharing/core/query"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

// Signin used to sign in user
func Signin(c echo.Context) error {
	return nil
}

// GetUsers get all users
func GetUsers(c echo.Context) error {
	users, err := query.GetUsers(middleware.Pagination{
		Page:  0,
		Limit: 50,
	})
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}
	_ = c.JSON(http.StatusOK, users)
	return nil
}

// GetUser get an user
func GetUser(c echo.Context) error {
	user, err := query.GetUserByID(database.User{
		ID: c.Param("id"),
	})
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}
	if user == nil {
		_ = c.JSON(http.StatusNotFound, "User not found")
		return nil
	}
	_ = c.JSON(http.StatusOK, user)
	return nil
}

// AddUser add an user
func AddUser(c echo.Context) error {
	user := database.User{}

	if err := c.Bind(&user); err != nil {
		fmt.Println(err)
		return err
	}

	if user.Mail == "" {
		_ = c.JSON(http.StatusBadRequest, "Mail can't be empty")
		return nil
	}
	if user.Login == "" {
		_ = c.JSON(http.StatusBadRequest, "Login can't be empty")
		return nil
	}
	if user.Username == "" {
		_ = c.JSON(http.StatusBadRequest, "Username can't be empty")
		return nil
	}
	if user.Password == "" {
		_ = c.JSON(http.StatusBadRequest, "Password can't be empty")
		return nil
	}
	if user.FirstName == "" {
		_ = c.JSON(http.StatusBadRequest, "Firtname can't be empty")
		return nil
	}
	if user.LastName == "" {
		_ = c.JSON(http.StatusBadRequest, "Lastname can't be empty")
		return nil
	}

	user.ID = helper.GenUlid()
	user.CreatedAt = time.Now()

	user.Password = helper.BcryptGen(user.Password)

	err := query.AddUser(user)
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}
	_ = c.JSON(http.StatusOK, user)
	return nil
}

// UpdateUser update an user
func UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser delete a user from our server
func DeleteUser(c echo.Context) error {
	err := query.DeleteUserByID(database.User{
		ID: c.QueryParam("id"),
	})
	if err != nil {
		_ = c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}
	_ = c.JSON(http.StatusOK, "User deleted")
	return nil
}
