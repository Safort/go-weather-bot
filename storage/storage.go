package storage

import (
    "../weathermap"
)


type StorageItem struct {
    weathermap.WeatherMap
    lastUpdate int
}

type Storage struct {
    cities map[string]StorageItem
}

func (self *Storage) Set(name string, info weathermap.WeatherMap) {
    self.cities[name] = StorageItem{info, 10}
}

func (self *Storage) Get(name string) StorageItem {
    return self.cities[name]
}


func NewStorage() *Storage {
    return &Storage{cities: make(map[string]StorageItem)}
}