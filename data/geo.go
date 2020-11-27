package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	data "weather/data/types"
)

// GetGeoData returns GeoData
func GetGeoData(date time.Time) data.GeoData {
	var structuredData data.GeoData

	geoKey := os.Getenv("GEO_KEY")
	url := fmt.Sprintf("https://api.ipgeolocation.io/astronomy?apiKey=%v&lat=%.4f&long=%.4f&date=%v", geoKey, Latitude, Longtitude, date.Format("2006-01-02"))

	headers := map[string]string{
		"accept": "application/json",
	}
	data := Get(url, "geo", headers)

	if err := json.Unmarshal(data, &structuredData); err != nil {
		log.Printf("JSON Error: %v\n %v", err, data)
	}

	return structuredData
}
