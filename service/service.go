package service

import (
	"io"
	"os"

	"github.com/OpenMeetingNotes/OpenMeetingNotes-server/model"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	validate = validator.New(validator.WithRequiredStructEnabled())
)

type media struct{}

type mediaUpload interface {
	FileRcv(c echo.Context) (model.File, error)
}

func NewMediaUpload() mediaUpload {
	return &media{}
}

func FileCopy(file model.File) (string, error) {

	// copy local
	dst, err := os.Create("data/" + file.Name)
	if err != nil {
		return "", err
	}

	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, file.File); err != nil {
		return "", err
	}

	return file.Name, nil
}

func (*media) FileRcv(c echo.Context) (model.File, error) {
	ret := model.File{}

	// Source
	file, err := c.FormFile("file")
	if err != nil {
		return ret, err
	}

	src, err := file.Open()
	if err != nil {
		return ret, err
	}
	defer src.Close()

	ret = model.File{
		Name: file.Filename,
		File: src,
	}

	// validate
	err = validate.Struct(file)
	if err != nil {
		return ret, err
	}

	return ret, nil
}
