package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Upload(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")

	fmt.Println(name, email)

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	dst, err := os.Create("data/" + file.Filename)
	if err != nil {
		return err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, fmt.Sprintf("File %s uploaded successfully", file.Filename))
}

func New(c echo.Context) error {
	// TODO: Implement handler.New
	return c.String(http.StatusOK, "Hello, World!")
}
