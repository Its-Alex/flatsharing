package middleware

import (
	"github.com/labstack/echo"
)

// Pagination struct to handle pagination
type Pagination struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}

// HydratePagination hydrate echo context with pagination
func HydratePagination(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var pagination Pagination

		_ = c.Bind(&pagination)

		if pagination.Page <= 0 {
			pagination.Page = 1
		}
		if pagination.Limit <= 0 {
			pagination.Limit = 50
		}

		c.Set("pagination", pagination)
		_ = next(c)
		return nil
	}
}
