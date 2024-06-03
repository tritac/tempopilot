package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/tritac/tempopilot/cmd/internals/appstore"
	worklog "github.com/tritac/tempopilot/cmd/internals/worklogs"
	api_services "github.com/tritac/tempopilot/cmd/services"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx       context.Context
	appStore  *appstore.AppStore
	apiClient *api_services.Client
}

// NewApp creates a new App application struct
func NewApp(appstore *appstore.AppStore, apiClient *api_services.Client) *App {
	return &App{appStore: appstore, apiClient: apiClient}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx

}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) onDomReady(ctx context.Context) {

	fmt.Println("Fom Ready")
}

func (a *App) CreateUserConfig(name, apiKey, userId string) (appstore.UserConfig, error) {

	res, err := a.appStore.StoreConfig(name, apiKey, userId, true)
	if err != nil {
		return appstore.UserConfig{}, nil
	}
	a.appStore.UserConfig.ApiKey = res.ApiKey
	a.appStore.UserConfig.UserId = res.UserId
	a.appStore.UserConfig.IsValidApi = true
	a.apiClient.ApiKey = res.ApiKey

	return res, nil
}

func (a *App) GetUserConfig() (appstore.UserConfig, error) {

	res, err := a.appStore.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res, "re")

	return a.appStore.LoadConfig()

}

func (a *App) VerifyApiKey(apiKey string) bool {

	req, err := http.NewRequest(http.MethodGet, host+"/accounts", nil)
	if err != nil {
		return false
	}
	req.Header.Add("Authorization", "Bearer "+apiKey)
	fmt.Println(apiKey)
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return false
	}
	defer resp.Body.Close()

	return resp.StatusCode == http.StatusOK

}

func (a *App) GetWorkLog(unixTime int64) ([]worklog.WorkLogResult, error) {
	date := time.UnixMilli(unixTime * 1000)
	fmt.Println(date)
	res, err := a.apiClient.GetUserBacklogByDate(date)
	if err != nil {
		return []worklog.WorkLogResult{}, err
	}
	return res, err
}

func (a *App) GetWorkLogTypes() ([]worklog.WorkLogAttr, error) {

	return a.apiClient.GetWorkLogAttribute()
}

type PostWorkLog struct {
	Key   string  `json:"key"`
	Value float64 `json:"value"`
}

func (a *App) PostWorkLog(work []PostWorkLog, date int64) {

	d := time.UnixMilli(date * 1000).Format("2006-01-02")

	createChan := make(chan worklog.WorkLogResult)
	createErr := make(chan error)
	for i, log := range work {

		attrs := make([]worklog.WorkLogValue, 0)
		arrt := worklog.WorkLogValue{Key: "_Test1_", Value: log.Key}
		attrs = append(attrs, arrt)

		l := worklog.CreateWorkLog{
			Attributes:       attrs,
			BillableSeconds:  log.Value * 60 * 60,
			WorkerID:         a.appStore.UserConfig.UserId,
			OriginTaskID:     "14914",
			AuthorAccountID:  a.appStore.UserConfig.UserId,
			TimeSpentSeconds: log.Value * 60 * 60,
			IssueID:          14914,
			StartDate:        d,
		}
		if l.BillableSeconds > 0 {
			go func() {
				fmt.Println("Channel initiated")
				res, err := a.apiClient.CreateWorkLog(l)
				if err != nil {
					createErr <- err

				}
				createChan <- res

			}()
			v, err := json.Marshal(l)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(v), i)
		}

	}

	select {
	case res := <-createChan:
		{
			runtime.EventsEmit(a.ctx, "worklog:created", res)
		}
	case err := <-createErr:
		{
			fmt.Println(err)
		}
	}

}

func (a *App) DeleteWorkflow(id int) (bool, error) {

	return a.apiClient.DeleteLog(id)
}
