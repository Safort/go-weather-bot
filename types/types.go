package types

import (
    "time"
)


type Coord struct {
    Lon, Lat float64
}

type Weather struct {
    Id int
    Main string
    Description string
    Icon string
}

type City struct {
    Id int
    Name string
    Coord Coord
    Country string
}

type Temp struct {
    Day, Min, Max, Night, Eve, Morn float64
}

type ListItem struct {
    Dt int
    Temp Temp
    Pressure float64
    Humidity int
    Weather []Weather
    Speed float64
    Deg float64
    Clouds int
}

type WeatherMap struct {
    City City
    Cod string
    List []ListItem
}

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


const (
    HELP = "/help"
    CITY = "/city"
)