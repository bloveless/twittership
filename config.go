package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type config struct {
	OAuthConsumerKey       string `json:"oauth_consumer_key"`
	OAuthConsumerSecret    string `json:"oauth_consumer_secret"`
	OAuthAccessToken       string `json:"oauth_access_token"`
	OAuthAccessTokenSecret string `json:"oauth_access_token_secret"`
}

func getConfigPath() (string, error) {
	xdgConfigHome := os.Getenv("XDG_CONFIG_HOME")
	if xdgConfigHome == "" {
		xdgConfigHome = os.Getenv("HOME") + "/.config"
	}

	configDirectory := xdgConfigHome + "/twittership"

	err := os.MkdirAll(configDirectory, 0755)
	if err != nil {
		return "", err
	}

	return configDirectory + "/config.json", nil
}

func loadConfig() (config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return config{}, err
	}

	configBytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		return config{}, err
	}

	c := config{}
	err = json.Unmarshal(configBytes, &c)
	if err != nil {
		return config{}, err
	}

	return c, nil
}

func saveConfig(c config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	bytes, err := json.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(configPath, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	_, err = f.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}
