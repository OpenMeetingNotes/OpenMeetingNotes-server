package handler

import (
	"fmt"
	"net/http"

	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/model"
	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/service"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func New(c echo.Context) error {
	formFile, err := service.NewMediaUpload().FileRcv(c)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.Response{
				Message: "Failed to receive file",
			},
		)
	}

	fileName, err := service.FileCopy(formFile)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			model.Response{
				Message: "Failed to create file",
			},
		)
	}

	return c.JSON(
		http.StatusOK,
		model.Response{
			Message: fmt.Sprintf("File %s uploaded successfully", fileName),
		},
	)
}
