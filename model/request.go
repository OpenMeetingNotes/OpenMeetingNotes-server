package model

import "mime/multipart"

type Request struct {
	SummaryLen string `json:"summaryLength"`
	Api        string `json:"api"`
	FilePath   string `json:"filePath"`
}

type Response struct {
	Message string `json:"message"`
}

type File struct {
	Name string         `json:"name,omitempty" validate:"required"`
	File multipart.File `json:"file,omitempty" validate:"required"`
}
