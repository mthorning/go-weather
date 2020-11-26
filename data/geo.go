package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

// GeoData from ipgeolacation API
type GeoData struct {
	Sunrise  string `json:"sunrise"`
	Sunset   string `json:"sunset"`
	Moonrise string `json:"moonrise"`
	Moonset  string `json:"moonset"`
}

// Fetch data from ipgeolocation.com
func Fetch() GeoData {
	var structuredData GeoData

	tomorrow := time.Now().Add(time.Hour * 24).Format("2006-01-02")

	geoKey := os.Getenv("GEO_KEY")
	url := fmt.Sprintf("https://api.ipgeolocation.io/astronomy?apiKey=%v&lat=%.4f&long=%.4f&date=%v", geoKey, Latitude, Longtitude, tomorrow)

	data := Get(url, "geo")

	if err := json.Unmarshal(data, &structuredData); err != nil {
		log.Printf("JSON Error: %v\n %v", err, data)
	}

	return structuredData
}
