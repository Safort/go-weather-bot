package storage

import (
    "testing"
    "../types"
)


func TestNewSet(t *testing.T) {
    store := NewStorage()

    city := types.WeatherMap {City: types.City{Name: "cityName"}}
    
    if res := store.Set("cityName", city); res.City.Name != "cityName" {
        t.Fail()
    }
}

func TestNewGet(t *testing.T) {
    store := NewStorage()

    city := types.WeatherMap {}

    store.Set("cityName", city)

    if _, ok := store.Get("cityName"); ok != true {
        t.Fail()
    }

    if _, ok := store.Get("cityName2"); ok == true {
        t.Fail()
    }
}
