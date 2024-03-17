package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func New(c echo.Context) error {
	// TODO: Implement handler.New
	return c.String(http.StatusOK, "Hello, World!")
}
