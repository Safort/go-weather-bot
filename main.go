package main

import (
    "log"
    "gopkg.in/telegram-bot-api.v4"
    "./config"
)

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

        log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

        msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

        bot.Send(msg)
    }
}
