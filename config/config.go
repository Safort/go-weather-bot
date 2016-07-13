package config

import (
    "io/ioutil"
    "encoding/json"
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