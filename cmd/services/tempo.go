package api_services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	worklog "github.com/tritac/tempopilot/cmd/internals/worklogs"
)

type WorkLogUIResponse struct {
	LogType string `json:"type"`
	Hours   int    `json:"hours"`
}

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

	return res.StatusCode == http.StatusOK

}

func (c *Client) GetUserBacklogByDate(t time.Time) ([]worklog.WorkLogResult, error) {

	date := t
	from := date.Format("2006-01-02")
	to := date.Format("2006-01-02")
	url := fmt.Sprintf("/worklogs?from=%s&to=%s", from, to)
	fmt.Println(url)
	res, err := c.do(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println(err)
		return []worklog.WorkLogResult{}, err
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		return []worklog.WorkLogResult{}, nil
	}
	var result worklog.WorkLogResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return []worklog.WorkLogResult{}, nil
	}

	return result.Results, nil

}
