package main

import (
	"fmt"
	"net/http"
	"time"
)

type (
	Client struct {
		host       string
		httpClient *http.Client
		apiKey     string
	}
)

func NewClient(host string, apiKey, apiSecret string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		host:       host,
		httpClient: client,
	}

}

func (c *Client) do(method, endpoint string, params map[string]string, apiKey string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+apiKey)
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}
