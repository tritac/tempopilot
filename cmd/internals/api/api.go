package api

import (
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	host       string
	httpClient *http.Client
	apiKey     string
}

type ApiClient struct {
	client *Client
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type successResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

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

func (c *ApiClient) Do(method, endpoint string, params map[string]string, apiKey string) (*http.Response, error) {
	baseURL := fmt.Sprintf("%s/%s", c.client.host, endpoint)
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
	return c.client.httpClient.Do(req)
}
