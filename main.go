package main

import (
	"fmt"
	"weather/data"

	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {
	geoData := data.Fetch()
	fmt.Println(geoData.Sunrise)
}
