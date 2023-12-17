package streamchecker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Category struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

type Stream struct {
	DaNick     string   `json:"daNick"`
	PreviewUrl string   `json:"previewUrl"`
	Category   Category `json:"category"`
	Title      string   `json:"title"`
	IsOnline   bool     `json:"isOnline"`
	IsEnded    bool     `json:"isEnded"`
}

type Getter struct {
	BaseUrl string
	Client  HTTPClient
}

func New(baseUrl string, client HTTPClient) *Getter {
	return &Getter{
		BaseUrl: baseUrl,
		Client:  client,
	}
}

func (c Getter) Get(nick string) (*Stream, error) {
	url := fmt.Sprintf(c.BaseUrl, nick)
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to build request: %w", err)
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to get API response: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response code is %v", resp.StatusCode)
	}

	var stream Stream
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&stream)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return &stream, nil
}
