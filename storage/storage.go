package storage

import (
    "time"
    "strconv"
    "../weathermap"
)


type StorageItem struct {
    weathermap.WeatherMap
    LastUpdate int
}

type Storage struct {
    cities map[string]StorageItem
}

func (self *Storage) Set(name string, info weathermap.WeatherMap) StorageItem {
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