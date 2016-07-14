package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "gopkg.in/telegram-bot-api.v4"
    "./config"
)


func Answer(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
    cnfg := config.Load()

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    url := "http://api.openweathermap.org/data/2.5/forecast/daily?"
    url += "APPID=" + cnfg.OwmToken + "&"
    url += "q=" + update.Message.Text + "&"
    url += "lang=ru"

    res, err := http.Get(url)
    if err != nil {
        log.Fatal(err)
    }
    json, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    if err != nil {
        log.Fatal(err)
    }

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(json))

    bot.Send(msg)
}



func main() {
    cnfg := config.Load()

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

        go Answer(bot, update)
    }
}
