package enotio

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"time"
)

type Config struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

const (
	DefaultBaseURL        = "https://api.enot.io"
	DefaultRequestTimeout = 30 * time.Second
)

type Client struct {
	httpClient *http.Client
	config     Config
}

func NewClient(cfg Config) *Client {
	client := &Client{
		config: cfg,
	}

	if client.config.BaseURL == "" {
		client.config.BaseURL = DefaultBaseURL
	}

	if client.config.HTTPClient != nil {
		client.httpClient = client.config.HTTPClient
	} else {
		client.httpClient = &http.Client{
			Timeout: DefaultRequestTimeout,
		}
	}

	return client
}

// sendRequest отправляет HTTP-запрос к API.
func (c *Client) sendRequest(method, url string, body []byte) ([]byte, error) {
	req, err := http.NewRequest(method, c.config.BaseURL+url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", c.config.APIKey)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("request failed with status: " + resp.Status)
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return responseBody, nil
}