package api_services

import (
	"encoding/json"
	"fmt"
	"io"
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
	res, err := c.do(http.MethodGet, "/accounts", nil, nil)
	if err != nil {
		return false
	}
	res.Header.Del("Authorization")
	res.Header.Add("Authorization", "Bearer "+apiKey)

	defer res.Body.Close()

	return res.StatusCode == http.StatusOK

}

func (c *Client) GetUserBacklogByDate(t time.Time) ([]worklog.WorkLogResult, bool, error) {

	date := t
	from := date.Format("2006-01-02")
	to := date.Format("2006-01-02")
	url := fmt.Sprintf("/worklogs?from=%s&to=%s", from, to)
	fmt.Println(c.ApiKey, "APID")

	res, err := c.do(http.MethodGet, url, nil, nil)

	if err != nil {
		return []worklog.WorkLogResult{}, true, err
	}
	if res.StatusCode == http.StatusUnauthorized {
		fmt.Println("ERORR VALID")
		return []worklog.WorkLogResult{}, false, err
	}

	body, err := io.ReadAll(res.Body) // response body is []byte

	if err != nil {
		return []worklog.WorkLogResult{}, true, nil
	}
	var result worklog.WorkLogResponse

	if err := json.Unmarshal(body, &result); err != nil {
		return []worklog.WorkLogResult{}, true, nil
	}

	return result.Results, true, nil

}

func (c *Client) GetWorkLogAttribute() ([]worklog.WorkLogAttr, error) {

	res, err := c.do(http.MethodGet, "/work-attributes", nil, nil)

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

func (c *Client) CreateWorkLog(workLog worklog.CreateWorkLog) (worklog.WorkLogResult, error) {

	payload, err := json.Marshal(workLog)
	if err != nil {
		fmt.Println(err)
	}
	res, err := c.do(http.MethodPost, "/worklogs", nil, payload)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	var result worklog.WorkLogResult
	if res.StatusCode == http.StatusOK {
		err = json.Unmarshal(body, &result)
		return result, err
	}
	return result, err

}

func (c *Client) DeleteLog(id int) (bool, error) {

	url := fmt.Sprintf("/worklogs/%d", id)
	fmt.Println(url)
	res, err := c.do(http.MethodDelete, url, nil, nil)
	if err != nil {
		fmt.Println(err, "Error")
	}
	return res.StatusCode == http.StatusOK || res.StatusCode == http.StatusNoContent, nil
}
