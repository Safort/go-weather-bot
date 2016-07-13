package main

import (
    "fmt"
    "./config"
)

func main() {
    cnfg := config.Load()

    fmt.Println("BotToken =", cnfg.BotToken)
    fmt.Println("OwmToken =", cnfg.OwmToken)
}