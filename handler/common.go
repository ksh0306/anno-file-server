package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Accessible(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Hello, Annotation AI customer.",
	})
}
