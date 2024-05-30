package main

import (
	"time"

	worklog "github.com/tritac/tempopilot/cmd/internals/worklogs"
)

func (a *App) GetMonthList(year int, month int) []worklog.WorkDay {

	m := time.Month(month)
	return a.appStore.WorkLog.GenerateMothDays(year, m)
}
