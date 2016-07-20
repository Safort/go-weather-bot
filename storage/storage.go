package storage

import (
    "time"
    "strconv"
    "../types"
)


type StorageItem struct {
    types.WeatherMap
    LastUpdate int
}

type Storage struct {
    cities map[string]StorageItem
}

func (self *Storage) Set(name string, info types.WeatherMap) StorageItem {
    t := time.Now()
    LastUpdate, _ := strconv.Atoi(t.Format("20060102150405"))
    self.cities[name] = StorageItem{info, LastUpdate}

    return self.cities[name]
}

func (self *Storage) Get(name string) (StorageItem, bool) {
    item, ok := self.cities[name]
    
    return item, ok
}

func NewStorage() *Storage {
    return &Storage{cities: make(map[string]StorageItem)}
}


func EnByRuDayMap(day string) time.Weekday {
    m := map[string]time.Weekday {
        "Понедельник": time.Monday,
        "Вторник": time.Tuesday,
        "Среда": time.Wednesday,
        "Четверг": time.Thursday,
        "Пятница": time.Friday,
        "Суббота": time.Saturday,
        "Воскресенье": time.Sunday,
    }

    return m[day]
}

func RuByEnDayMap(day time.Weekday) string {
    m := map[time.Weekday]string {
        time.Monday: "Понедельник",
        time.Tuesday: "Вторник",
        time.Wednesday: "Среда",
        time.Thursday: "Четверг",
        time.Friday: "Пятница",
        time.Saturday: "Суббота",
        time.Sunday: "Воскресенье",
    }

    return m[day]
}
