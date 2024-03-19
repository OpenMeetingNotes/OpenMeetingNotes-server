package main

import (
	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/hello", handler.Hello)
	e.POST("/new", handler.New)

	e.Logger.Fatal(e.Start("localhost:8090"))
}
