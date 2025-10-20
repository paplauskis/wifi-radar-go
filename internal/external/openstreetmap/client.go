package openstreetmap

import (
	"io"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	ApiURL     string
}

func NewClient(city string) *Client {
	return &Client{
		httpClient: &http.Client{},
		ApiURL:     "https://nominatim.openstreetmap.org/search?q=" + city + "&format=json&limit=1",
	}
}

func (c *Client) Query() ([]byte, error) {
	req, err := http.NewRequest("GET", c.ApiURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "wifi-radar-go")

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
