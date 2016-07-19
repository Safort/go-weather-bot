package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "bytes"
    "encoding/json"
    "strings"
    "time"
    "strconv"
    "gopkg.in/telegram-bot-api.v4"
    "./config"
    "./weathermap"
    "./storage"
)


type CityInfo struct {
    Name string
    Day time.Weekday
    WindSpeed string
    Clouds string
    TempMorn string
    TempDay string
    TempEve string
    TempNight string
}


func requestCityInfo(cityName string) weathermap.WeatherMap {
    var weather weathermap.WeatherMap
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


func commandHelp() string {
    return "go-weather-bot - это всего лишь попытка прокачать себя в новой" +
           " области программирования на языке Go."
}


func formCityInfo(city CityInfo) string {
    return "Город: " + city.Name + "\n" +
           "День: " + storage.RuByEnDayMap(city.Day) + "\n" +
           "Скорость ветра: " + city.WindSpeed + " м/с \n" +
           "Облачность: " + city.Clouds + "%\n" +
           "Температура\n" +
           "- Утро: " + city.TempMorn + "°C\n" +
           "- День: " + city.TempDay + "°C\n" +
           "- Вечер: " + city.TempEve + "°C\n" +
           "- Ночь: " + city.TempNight + "°C\n"
}


func commandCity(city storage.StorageItem) string {
    var cityInfo CityInfo
    cityInfo.Name = city.City.Name
    presentDay := time.Now().Weekday()

    for i, list := 0, city.List; i < len(list); i++ {
        item := list[i]
        weekday := time.Unix(int64(item.Dt), 0).Weekday()

        if weekday == presentDay {
            cityInfo.Day = weekday
            cityInfo.WindSpeed = strconv.FormatFloat(item.Speed, 'f', -1, 64)
            cityInfo.Clouds = strconv.Itoa(item.Clouds)

            cityInfo.TempMorn = strconv.FormatFloat(item.Temp.Morn, 'f', -1, 64)
            cityInfo.TempDay = strconv.FormatFloat(item.Temp.Day, 'f', -1, 64)
            cityInfo.TempEve = strconv.FormatFloat(item.Temp.Eve, 'f', -1, 64)
            cityInfo.TempNight = strconv.FormatFloat(item.Temp.Night, 'f', -1, 64)

            break
        }
    }

    return formCityInfo(cityInfo)
}


func main() {
    cnfg  := config.Load()
    store := storage.NewStorage()
    bot, err := tgbotapi.NewBotAPI(cnfg.BotToken)

    const (
        HELP = "/help"
        CITY = "/city"
    )

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

        var answerText, cityName string
        args := strings.Split(update.Message.Text, " ")
        command := args[0]

        if len(args) > 1 {
            cityName = args[1]
        }

        switch command {
            case HELP:
                answerText = commandHelp()
            
            case CITY:
                answerText = "cityName == " + cityName
                var city storage.StorageItem
                var ok bool

                if city, ok = store.Get(cityName); ok == false {
                    city = store.Set(cityName, requestCityInfo(cityName));
                }

                if len(args) == 2 {
                    answerText = commandCity(city) //Погода на сегодня
                } else if len(args) == 3 {
                    answerText = "Погода на всю неделю или выбранный день"
                }
            
            default:
                answerText = "404. Такой команды не существует."
        }

        store.Get("")
        msg := tgbotapi.NewMessage(update.Message.Chat.ID, answerText)
        bot.Send(msg)
    }
}
