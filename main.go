package main

import (
	"embed"
	"time"

	"github.com/tritac/tempopilot/cmd/internals/appstore"
	api_services "github.com/tritac/tempopilot/cmd/services"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

const (
	host = "https://api.tempo.io/4"
)

func main() {
	// Create an instance of the app structure
	appStore := appstore.NewAppStore()

	apiClient := api_services.NewClient(host, appStore.UserConfig.ApiKey, time.Second*5)

	app := NewApp(appStore, apiClient)

	// Create application with options

	err := wails.Run(&options.App{
		Title:     "Tempo Pilot",
		Width:     750,
		Height:    500,
		MinWidth:  600,
		MinHeight: 400,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       app.onDomReady,

		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
