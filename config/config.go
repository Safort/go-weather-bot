package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	BotToken string
	OwmToken string
}

func Load() Config {
	var config Config

	res, _ := ioutil.ReadFile("./config.json")
	jsonBlob := []byte(string(res))

	json.Unmarshal(jsonBlob, &config)

	return config
}
