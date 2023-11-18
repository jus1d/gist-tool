package gist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/jus1d/gist-tool/internal/file"
	"io"
	"net/http"
)

const url = "https://api.github.com/gists"

type Conn struct {
	token string
}

func New(token string) *Conn {
	return &Conn{token: token}
}

func (c *Conn) Create(filepath string) (string, error) {
	filename, content, err := file.Read(filepath)
	if err != nil {
		return "", ErrFileNotExists
	}

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
