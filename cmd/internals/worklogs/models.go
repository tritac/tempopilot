package worklog

import "time"

type MetaData struct {
	Count  int `json:"count"`
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

type Issue struct {
	Self string `json:"self"`
	ID   int    `json:"id"`
}

type WorkLogResult struct {
	Self             string           `json:"self"`
	TempoWorklogID   int              `json:"tempoWorklogId"`
	Issue            Issue            `json:"issue"`
	TimeSpentSeconds int              `json:"timeSpentSeconds"`
	BillableSeconds  int              `json:"billableSeconds"`
	StartDate        string           `json:"startDate"`
	StartTime        string           `json:"startTime"`
	Description      string           `json:"description"`
	CreatedAt        time.Time        `json:"createdAt"`
	UpdatedAt        time.Time        `json:"updatedAt"`
	Attributes       WorkLogAttribute `json:"attributes"`
}

type WorklogAuthor struct {
	Self      string `json:"self"`
	AccountID string `json:"accountId"`
}

type WorkLogValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type WorkLogAttribute struct {
	Self   string         `json:"self"`
	Values []WorkLogValue `json:"values"`
}

type WorkLogResponse struct {
	Self     string          `json:"self"`
	Metadata MetaData        `json:"metadata"`
	Results  []WorkLogResult `json:"results"`
}
