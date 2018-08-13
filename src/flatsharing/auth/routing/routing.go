package routing

import (
	"flatsharing/auth/controllers"

	"github.com/labstack/echo"
)

// Setup api routeur
func Setup(e *echo.Echo) {
	// e.Use(middleware.HydratePagination)

	v1(e.Group("/v1"))
}

func v1(e *echo.Group) {
	e.POST("/singin", controllers.Signin)

	e.GET("/users", controllers.GetUsers)
	e.GET("/user/:id", controllers.GetUser)

	e.POST("/user", controllers.AddUser)
	e.PUT("/user/:id", controllers.UpdateUser)
	e.DELETE("/user/:id", controllers.DeleteUser)
}
