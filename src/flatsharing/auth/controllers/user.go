package controllers

import (
	"flatsharing/core/database"
	"flatsharing/core/helper"
	"flatsharing/core/middleware"
	"flatsharing/core/query"
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

	if err := c.Bind(&fields); err != nil {
		return err
	}

	err := helper.Validate.Struct(fields)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return err
	}

	user, err := query.GetUserByMail(database.User{
		Mail: fields.Mail,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}
	if user == nil {
		c.JSON(http.StatusNotFound, "User not found")
		return err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(fields.Password)) != nil {
		c.JSON(http.StatusUnauthorized, "Password mismatch!")
		return err
	}

	token, err := query.CreateToken(*user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Internal server error")
		return err
	}

	c.JSON(http.StatusOK, map[string]string{
		"token": token.Token,
	})
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
		return err
	}

	err := helper.Validate.Struct(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return nil
	}

	user.ID = helper.GenUlid()

	user.Password = helper.BcryptGen(user.Password)

	err = query.AddUser(user)
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
