package main

import (
	"net/http"

	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/upload", handler.Upload)
	e.POST("/process", handler.New)

	e.Logger.Fatal(e.Start(":8090"))
}
