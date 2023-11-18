package gist

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const url = "https://api.github.com/gists"

type Conn struct {
	token string
}

func New(token string) *Conn {
	return &Conn{token: token}
}

func (c *Conn) Create(filename string) (string, error) {
	contentBytes, err := os.ReadFile(filename)
	if err != nil {
		return "", ErrFileNotExists
	}
	content := string(contentBytes)

	body := CreateRequest{
		Description: "",
		Public:      true,
		Files: map[string]File{
			filename: {Content: content},
		},
	}

	jsonData, err := json.Marshal(body)
	if err != nil {
		return "", ErrInvalidRequest
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", ErrInvalidRequest
	}

	c.addHeaders(req)

	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return "", ErrInvalidRequest
	}
	defer response.Body.Close()

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", ErrCantReadBody
	}

	var resp CreateResponse

	err = json.Unmarshal(responseBody, &resp)
	if err != nil {
		return "", err
	}

	return resp.URL, nil
}

func (c *Conn) addHeaders(r *http.Request) {
	r.Header.Set("Accept", "application/vnd.github+json")
	r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	r.Header.Set("X-GitHub-Api-Version", "2022-11-28")
}

type CreateRequest struct {
	Description string          `json:"description"`
	Public      bool            `json:"public"`
	Files       map[string]File `json:"files"`
}

type CreateResponse struct {
	URL string `json:"html_url"`
}

type File struct {
	Content string `json:"content"`
}

var (
	ErrInvalidRequest = errors.New("can't create a request")
	ErrFileNotExists  = errors.New("file doesn't exists")
	ErrCantReadBody   = errors.New("can't read response body")
)
