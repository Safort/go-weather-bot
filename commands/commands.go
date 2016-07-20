package commands

import (
    "strconv"
    "time"
    "../storage"
    "../types"
)


func formCityInfo(city types.CityInfo) string {
    return "День: " + storage.RuByEnDayMap(city.Day) + "\n" +
           "Скорость ветра: " + city.WindSpeed + " м/с \n" +
           "Облачность: " + city.Clouds + "%\n" +
           "Температура\n" +
           "- Утро: " + city.TempMorn + "°C\n" +
           "- День: " + city.TempDay + "°C\n" +
           "- Вечер: " + city.TempEve + "°C\n" +
           "- Ночь: " + city.TempNight + "°C\n"
}


func CommandHelp() string {
    return "go-weather-bot - это всего лишь попытка прокачать себя в новой" +
           " области программирования на языке Go."
}

func CommandCity(city storage.StorageItem, day time.Weekday) string {
    var cityInfo types.CityInfo
    cityInfo.Name = city.City.Name

    for i, list := 0, city.List; i < len(list); i++ {
        item := list[i]
        weekday := time.Unix(int64(item.Dt), 0).Weekday()

        if weekday == day {
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

    return "Город: " + cityInfo.Name + "\n" + formCityInfo(cityInfo)
}


func CommandCityFullWeek(city storage.StorageItem) string {
    var info string

    for i, list := 0, city.List; i < len(list); i++ {
        var cityInfo types.CityInfo
        item := list[i]
        weekday := time.Unix(int64(item.Dt), 0).Weekday()

        cityInfo.Day = weekday
        cityInfo.WindSpeed = strconv.FormatFloat(item.Speed, 'f', -1, 64)
        cityInfo.Clouds = strconv.Itoa(item.Clouds)

        cityInfo.TempMorn = strconv.FormatFloat(item.Temp.Morn, 'f', -1, 64)
        cityInfo.TempDay = strconv.FormatFloat(item.Temp.Day, 'f', -1, 64)
        cityInfo.TempEve = strconv.FormatFloat(item.Temp.Eve, 'f', -1, 64)
        cityInfo.TempNight = strconv.FormatFloat(item.Temp.Night, 'f', -1, 64)

        info += formCityInfo(cityInfo) + "\n---\n"
    }

    return "Город: " + city.City.Name + "\n\n" + info
}