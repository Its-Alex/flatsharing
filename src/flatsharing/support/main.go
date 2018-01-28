package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type request struct {
	Platform    string `json:"platform"`
	Description string `json:"description"`
}

type response struct {
	Msg string `json:"msg"`
}

func createIssue(c echo.Context) error {
	var req request
	c.Bind(&req)
	fmt.Println(req)
	return c.JSON(http.StatusCreated, response{
		Msg: "Ressource has been created",
	})
}

func main() {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/", createIssue)

	e.Logger.Fatal(e.Start(":1323"))
}
