package api_services

import (
	"encoding/json"
	"fmt"
	"io"
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
	fmt.Println(c.ApiKey, "APID")
	res, err := c.do(http.MethodGet, url, nil)
	if err != nil {
		return []worklog.WorkLogResult{}, err
	}
	body, err := ioutil.ReadAll(res.Body) // response body is []byte
	if err != nil {
		return []worklog.WorkLogResult{}, nil
	}
	var result worklog.WorkLogResponse
	fmt.Println(string(body))
	if err := json.Unmarshal(body, &result); err != nil {
		return []worklog.WorkLogResult{}, nil
	}

	return result.Results, nil

}

func (c *Client) GetWorkLogAttribute() ([]worklog.WorkLogAttr, error) {

	res, err := c.do(http.MethodGet, "/work-attributes", nil)

	if err != nil {
		return nil, err
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var result worklog.WorkLogTypeResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println(err)
	}
	return result.Results, nil

}
