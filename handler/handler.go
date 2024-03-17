package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"fmt"
	"bytes"
	"mime/multipart"
	"os"
	"path/filepath"
	"io"
	"strings"
)


func New(c echo.Context) error {

	// Inputs: filepath(of .mp3) (hardcode for now), length of summary, model type(default ChatGPT)
	
	// Call apis
	// Set up secrets locally?


  // Whisper API 
  url := "https://api.openai.com/v1/audio/transcriptions"
  method := "POST"

  payload := &bytes.Buffer{}
  writer := multipart.NewWriter(payload)
  file, errFile1 := os.Open("/path/to/file")
  defer file.Close()
  part1,
         errFile1 := writer.CreateFormFile("file",filepath.Base("/path/to/file"))
  _, errFile1 = io.Copy(part1, file)
  if errFile1 != nil {
    fmt.Println(errFile1)
    return
  }
  _ = writer.WriteField("model", "whisper-1")
  err := writer.Close()
  if err != nil {
    fmt.Println(err)
    return
  }

  client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
	fmt.Println(err)
	return
	}

	// TODO: Get api key and put it in a secret --> secret.env
	openapi_key := ""
	bearer_token := fmt.Sprintf("Bearer %s", openapi_key)

	req.Header.Add("Authorization", bearer_token)

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
	fmt.Println(err)
	return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
	fmt.Println(err)
	return
	}

	// End of Whisper API


	// ChatGPT API


	url := "https://api.openai.com/v1/chat/completions"
	method := "POST"

	model_type := "gpt-3.5-turbo"
	word_count := 250
	content := string(body)

	chat_request := fmt.Stringf(`{`+"
	"+`
		"model": %s,`+"
	"+`
		"messages": [`+"
	"+`
		{`+"
	"+`
			"role": "You are a writing assistant that summarizes text conversations.",`+"
	"+`
			"content": "You will be given the following text conversation after the colon. Summarize the text conversation within %s words so that it includes the main points of the discussion: %s"`+"
	"+`
	}`, model_type,word_count, content)


	payload := strings.NewReader()

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", bearer_token)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	return 
	fmt.Println(string(body))
	}

	// End of ChatGPT API

