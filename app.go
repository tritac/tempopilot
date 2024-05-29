package main

import (
	"context"
	"fmt"

	"github.com/tritac/tempopilot/cmd/internals/appstore"
)

// App struct
type App struct {
	ctx      context.Context
	appStore *appstore.AppStore
}

// NewApp creates a new App application struct
func NewApp(appstore *appstore.AppStore) *App {
	return &App{appStore: appstore}
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

func (a *App) CreateUserConfig(name, apiKey, userId string) {
	a.appStore.StoreConfig(name, apiKey, userId)
}

func (a *App) GetUserConfig() (appstore.UserConfig, error) {

	return a.appStore.LoadConfig()

}
