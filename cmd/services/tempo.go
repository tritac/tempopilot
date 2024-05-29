package api_services

import (
	"fmt"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func (c *Client) VerifyApi(apiKey string) bool {
	res, err := c.do(http.MethodGet, "/accounts", nil)
	if err != nil {
		return false
	}
	res.Header.Del("Authorization")
	res.Header.Add("Authorization", "Bearer "+apiKey)

	defer res.Body.Close()
	fmt.Println(res.StatusCode)
	if res.StatusCode == http.StatusOK {
		return true
	}
	return false

}
