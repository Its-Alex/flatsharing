package main

import (
	"bytes"
	"net/http"
	"os"

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

	url := "https://api.github.com/repos/flatsharing/" + req.Platform + "/issues"
	data := []byte(`{"title":"Automatic issue from bot", "body": "` + req.Description + `"}`)

	call, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	call.Header.Set("Content-Type", "application/json")
	call.Header.Set("Authorization", "token "+os.Getenv("GITHUB_TOKEN"))
	client := http.Client{}
	res, err := client.Do(call)
	if err != nil {
		return c.JSON(http.StatusConflict, response{
			Msg: "Request error",
		})
	}
	defer res.Body.Close()

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
