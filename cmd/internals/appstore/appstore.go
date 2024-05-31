package appstore

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	worklog "github.com/tritac/tempopilot/cmd/internals/worklogs"
)

type UserConfig struct {
	ApiKey     string `json:"api_key"`
	UserName   string `json:"user_name"`
	UserId     string `json:"user_id"`
	IsValidApi bool   `json:"isValidApi"`
}
type AppStore struct {
	UserConfig UserConfig
	ConfigPath string
	WorkLog    *worklog.WorkLogStore
}

func NewAppStore() *AppStore {

	configFilePath, err := xdg.ConfigFile("tempopilot/config.json")

	if err != nil {
		panic(err)
	}

	_, err = os.Stat(configFilePath)
	if err != nil {
		// check if not file error
		fileExists := os.IsExist(err)
		// if file not present , create the file, with empty content
		if !fileExists {
			file, err := os.OpenFile(configFilePath, os.O_RDWR|os.O_CREATE, 0644)
			if err != nil {
				panic(err)
			}

			// if file dosnt exits , create a blank file
			if _, err := file.WriteString(`{"api_key":""}`); err != nil {
				panic(err)
			}
			defer file.Close()
		}
	}
	content, err := os.ReadFile(configFilePath)

	if err != nil {
		panic(err)
	}

	var config UserConfig

	err = json.Unmarshal(content, &config)

	if err != nil {
		panic(err)
	}

	workLog := worklog.NewWorkLogStore()

	return &AppStore{ConfigPath: configFilePath, WorkLog: workLog, UserConfig: config}
}

func (as *AppStore) StoreConfig(name, apiKey, userId string, isVerified bool) (UserConfig, error) {
	userConfig := UserConfig{ApiKey: apiKey, UserName: name, UserId: userId, IsValidApi: isVerified}
	userJson, err := json.Marshal(userConfig)
	if err != nil {
		fmt.Println(err)
	}
	// make sure to open in truncate mode, O_TRUNC writes from 0,0
	fi, err := os.OpenFile(as.ConfigPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return UserConfig{}, err
	}
	defer fi.Close()

	content, err := fi.Write(userJson)
	if content != 0 {
		userConf, err := as.LoadConfig()

		if err != nil {
			return userConf, err
		}
		return userConf, err
	}

	return UserConfig{}, err

}

func (as *AppStore) LoadConfig() (UserConfig, error) {
	confFile, err := os.OpenFile(as.ConfigPath, os.O_RDONLY, 0644)
	if err != nil {
		return UserConfig{}, err
	}
	defer confFile.Close()

	var userConfig UserConfig
	err = json.NewDecoder(confFile).Decode(&userConfig)
	if err != nil {
		return UserConfig{}, err
	}
	return userConfig, nil
}
