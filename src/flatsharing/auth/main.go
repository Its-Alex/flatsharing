package main

import (
	"flatsharing/auth/routing"
	"flatsharing/core/database"
	"flatsharing/core/helper"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	database.Ping()
}

func echoSetup(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.BodyLimit("2M"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderAuthorization, echo.HeaderContentType},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
}

func main() {
	e := echo.New()
	echoSetup(e)

	routing.Setup(e)

	e.Logger.Fatal(e.Start(helper.GetEnv("AUTH_PORT", ":3000")))
}
