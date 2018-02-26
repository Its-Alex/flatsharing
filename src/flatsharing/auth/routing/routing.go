package routing

import (
	"flatsharing/auth/controllers"
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
	e.GET("/users", controllers.GetUsers)
	e.POST("/user", controllers.AddUser)
	e.GET("/user/:id", controllers.GetUser)
	e.PUT("/user/:id", controllers.UpdateUser)
}
