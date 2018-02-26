package routing

import (
	"flatsharing/core/middleware"

	"github.com/labstack/echo"
)

// Setup api routeur
func Setup(e *echo.Echo) {
	e.Use(middleware.HydratePagination)
	e.Use(middleware.Reponse)

	v1(e.Group("/v1"))
}

func v1(e *echo.Group) {
	e.GET("/", func(c echo.Context) error {
		_ = c.JSON(200, "Hello, world!")
		return nil
	})
}
