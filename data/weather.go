package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	data "weather/data/types"
)

// GetWeatherData returns WeatherData
func GetWeatherData() data.WeatherData {
	var structuredData data.WeatherData

	metID := os.Getenv("MET_ID")
	metSecret := os.Getenv("MET_SECRET")
	url := fmt.Sprintf("https://api-metoffice.apiconnect.ibmcloud.com/metoffice/production/v0/forecasts/point/hourly?latitude=%v&longitude=%v", Latitude, Longtitude)

	headers := map[string]string{
		"accept":              "application/json",
		"x-ibm-client-secret": metSecret,
		"x-ibm-client-id":     metID,
	}
	data := Get(url, "weather", headers)

	if err := json.Unmarshal(data, &structuredData); err != nil {
		log.Printf("JSON Error: %v\n %v", err, data)
	}

	return structuredData
}
