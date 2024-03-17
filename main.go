package main

import (
	"net/http"

	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/new", handler.New)
	e.Logger.Fatal(e.Start(":8090"))
}
