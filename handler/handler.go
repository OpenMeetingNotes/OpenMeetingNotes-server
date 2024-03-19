package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
)

type Chat_Request struct {
	Model    string   `json:"model"`
	Messages Messages `json:"messages"`
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func Config() string {
	secret := os.Getenv("OPENAI_SECRET")
	return secret
}

func New(c echo.Context) error {

	// Inputs: filepath(of .mp3) (hardcode for now), length of summary, model type(default ChatGPT)

	// Whisper API
	url := "https://api.openai.com/v1/audio/transcriptions"
	method := "POST"

	w_payload := &bytes.Buffer{}
	writer := multipart.NewWriter(w_payload)

	file, errFile1 := os.Open("./data/test.mp3")

	if errFile1 != nil {
		fmt.Println(errFile1)
		defer file.Close()
		return c.String(http.StatusTeapot, "Test Error")
	}

	defer file.Close()

	part1, errFile1 := writer.CreateFormFile("file", filepath.Base("/path/to/file"))
	if errFile1 != nil {
		fmt.Println(errFile1)
		return c.String(http.StatusTeapot, "Test Error")
	}

	_, errFile1 = io.Copy(part1, file)
	if errFile1 != nil {
		fmt.Println(errFile1)
		return c.String(http.StatusTeapot, "Test Error")
	}

	_ = writer.WriteField("model", "whisper-1")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, w_payload)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	openapi_key := ""
	bearer_token := fmt.Sprintf("Bearer %s", openapi_key)

	req.Header.Add("Authorization", bearer_token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	// End of Whisper API

	// ChatGPT API

	url = "https://api.openai.com/v1/chat/completions"
	method = "POST"

	model_type := "gpt-3.5-turbo"
	word_count := 250
	content := fmt.Sprintf("You will be given the following text conversation after the colon. Summarize the text conversation within %d words so that it includes the main points of the discussion: %s", word_count, string(body))

	request_object := Chat_Request{
		Model: model_type,
		Messages: Messages{
			Role:    "You are a writing assistant that summarizes text conversations.",
			Content: content,
		},
	}

	jsonData, err := json.Marshal(request_object)

	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return c.String(http.StatusTeapot, "Test error")
	}

	requ := string(jsonData)
	payload := strings.NewReader(requ)

	client = &http.Client{}
	req, err = http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	res, err = client.Do(req)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	defer res.Body.Close()

	body, err = io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusTeapot, "Test Error")
	}

	return c.String(http.StatusOK, string(body))
	// End of ChatGPT API
}
