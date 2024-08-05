package middleware

import (
	"backend-github/model"
	"github.com/labstack/echo"
	"net/http"
)

func IsAdmin() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role := c.Get("role").(string)
			if role != "ADMIN" {
				return c.JSON(http.StatusForbidden,
					model.Response{
						StatusCode: http.StatusForbidden,
						Message:    "You are not authorized to access this resource",
						Data:       nil,
					})
			}
			return next(c)
		}
	}
}
