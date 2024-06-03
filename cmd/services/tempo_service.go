package api_services

import (
	"bytes"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	host       string
	httpClient *http.Client
	ApiKey     string
}

func NewClient(host string, apiKey string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		host:       host,
		httpClient: client,
		ApiKey:     apiKey,
	}
}

func (c *Client) do(method, endpoint string, params map[string]string, body []byte) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+c.ApiKey)

	q := req.URL.Query()

	for key, val := range params {
		q.Set(key, val)
	}

	req.URL.RawQuery = q.Encode()

	return c.httpClient.Do(req)
}
