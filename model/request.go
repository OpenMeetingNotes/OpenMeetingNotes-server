package model

type Request struct {
	SummaryLen string `json:"summaryLength"`
	Api        string `json:"api"`
	FilePath   string `json:"filePath"`
}
