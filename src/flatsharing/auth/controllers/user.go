package controllers

import (
	"flatsharing/core/database"
	"flatsharing/core/helper"
	"flatsharing/core/middleware"
	"flatsharing/core/query"
	"fmt"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/labstack/echo"
)

// SigninFields contain needed info to sign in
type SigninFields struct {
	Mail     string `json:"mail" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8,lte=64"`
}

// Signin used to sign in user
func Signin(c echo.Context) error {
	fields := SigninFields{}

	// Check fields
	if err := c.Bind(&fields); err != nil {
		return err
	}
	if err := helper.Validate.Struct(fields); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	// Get user
	user, err := query.GetUserByMailOrID(database.User{
		Mail: fields.Mail,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(fields.Password)) != nil {
		return c.JSON(http.StatusUnauthorized, "Password mismatch!")
	}

	token, err := query.CreateToken(*user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token.Token,
	})
}

// GetUsers get all users
func GetUsers(c echo.Context) error {
	users, err := query.GetUsers(middleware.Pagination{
		Page:  0,
		Limit: 50,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, users)
}

// GetUser get an user
func GetUser(c echo.Context) error {
	user, err := query.GetUserByMailOrID(database.User{
		ID:   c.Param("id"),
		Mail: c.Param("id"),
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "User not found")
	}
	return c.JSON(http.StatusOK, user)
}

// AddUser add an user
func AddUser(c echo.Context) error {
	user := database.User{}

	if err := c.Bind(&user); err != nil {
		return err
	}
	err := helper.Validate.Struct(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	user.ID = helper.GenUlid()
	user.Password = helper.BcryptGen(user.Password)

	err = query.AddUser(user)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, user)
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
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}
	return c.JSON(http.StatusOK, "User deleted")
}
