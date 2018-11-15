package middleware

import (
	"github.com/Its-Alex/flatsharing/internal/core/helper"

	"github.com/labstack/echo"
)

// Reponse middleware to handle response
func Reponse(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		_ = next(ctx)

		if err, ok := ctx.Get("return_error").(helper.Error); ok {
			if err.Type != "" && err.Throwable {
				code, ok := helper.ErrorsHTTP[err.Type]
				if !ok {
					code = 500
				}

				_ = ctx.JSON(code, err)
			} else {
				data := ctx.Get("return")
				returnCode := ctx.Get("return_code")
				code, _ := returnCode.(int)
				_ = ctx.JSON(code, data)
			}
		} else {
			_ = ctx.JSON(404, helper.Error{
				Message: helper.BuildErrorMessage(404, "route"),
				Type:    helper.UnknownResource,
			})
		}
		return nil
	}
}
