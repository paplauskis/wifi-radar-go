package overpass

import (
	"bytes"
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	ApiURL     string
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
		ApiURL:     "https://overpass-api.de/api/interpreter",
	}
}

func (c *Client) Query(query string) ([]byte, error) {
	resp, err := c.httpClient.Post(
		c.ApiURL,
		"application/x-www-form-urlencoded; charset=UTF-8",
		bytes.NewBufferString(query))

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
