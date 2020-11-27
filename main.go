package main

import (
	"fmt"
	"time"
	"weather/data"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	tomorrow := time.Now().Add(time.Hour * 24)
	geoData := data.GetGeoData(tomorrow)
	weather := data.GetWeatherData()
	tides := data.GetTidesData()
	fmt.Println(geoData.Sunrise, weather.Type, tides)
}
