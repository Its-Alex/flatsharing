package middleware

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
)

// Pagination struct to handle pagination
type Pagination struct {
	Page  int
	Limit int
}

// HydratePagination hydrate echo context with pagination
func HydratePagination(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		pagination := Pagination{
			Page:  1,
			Limit: 50,
		}

		if val := c.QueryParam("page"); val != "" {
			page, _ := strconv.Atoi(val)
			if page > 0 {
				pagination.Page = page
			}
		}
		if val := c.QueryParam("limit"); val != "" {
			limit, _ := strconv.Atoi(val)
			if limit > 0 {
				pagination.Limit = limit
			}
		}

		fmt.Println(pagination)
		c.Set("pagination", pagination)
		next(c)
		return nil
	}
}
