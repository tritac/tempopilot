package api_services

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
		apiSecret  string
	}
)

func NewClient(host string, apiKey string, timeout time.Duration) *Client {
	client := &http.Client{
		Timeout: timeout,
	}
	return &Client{
		host:       host,
		httpClient: client,
		apiKey:     apiKey,
	}
}

func (c *Client) do(method, endpoint string, params map[string]string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.host, endpoint)
	req, err := http.NewRequest(method, baseURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	fmt.Println(c.apiKey, "-----------------------c-c-cc-c-c-")
	req.Header.Add("Authorization", "Bearer "+c.apiKey)
	q := req.URL.Query()
	for key, val := range params {
		q.Set(key, val)
	}
	req.URL.RawQuery = q.Encode()
	return c.httpClient.Do(req)
}
