package main

import (
	"./commands"
	"./config"
	"./storage"
	"./types"
	"bytes"
	"encoding/json"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func requestCityInfo(cityName string) types.WeatherMap {
	var weather types.WeatherMap
	cnfg := config.Load()

	// Create request url
	url := "http://api.openweathermap.org/data/2.5/forecast/daily?"
	url += "APPID=" + cnfg.OwmToken + "&"
	url += "q=" + cityName + "&"
	url += "lang=ru&units=metric"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	resText, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	dec := json.NewDecoder(bytes.NewReader(resText))
	dec.Decode(&weather)

	return weather
}

func sendAnswer(bot *tgbotapi.BotAPI, update tgbotapi.Update, store *storage.Storage) {
	var answerText, cityName string
	args := strings.Split(update.Message.Text, " ")
	command := args[0]

	log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

	if len(args) > 1 {
		cityName = args[1]
	}

	switch command {
	case types.HELP:
		answerText = commands.CommandHelp()

	case types.CITY:
		var (
			city storage.StorageItem
			ok   bool
		)
		t := time.Now()
		currentTime, _ := strconv.Atoi(t.Format("20060102150405"))

		if city, ok = store.Get(cityName); ok == false {
			city = store.Set(cityName, requestCityInfo(cityName))
		} else {
			if (currentTime - city.LastUpdate) > 60 {
				city = store.Set(cityName, requestCityInfo(cityName))
			}
		}

		if len(args) == 2 {
			answerText = commands.CommandCity(city, time.Now().Weekday())
		} else if len(args) == 3 && args[2] == "week" {
			answerText = commands.CommandCityFullWeek(city)
		} else {
			answerText = commands.CommandCity(city, storage.EnByRuDayMap(args[2]))
		}

	default:
		answerText = commands.CommandNotFound()
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
	bot.Send(msg)
}

func main() {
	cnfg := config.Load()
	store := storage.NewStorage()
	bot, err := tgbotapi.NewBotAPI(cnfg.BotToken)

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		go sendAnswer(bot, update, store)
	}
}
