package data

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	data "weather/data/types"
)

// GetTidesData returns TidesData
func GetTidesData() data.TidesData {
	var structuredData data.TidesData

	tideKey := os.Getenv("TIDE_KEY")
	url := fmt.Sprintf("https://tides.p.rapidapi.com/tides?latitude=%v&longitude=%v", Latitude, Longtitude)

	headers := map[string]string{
		"accept":          "application/json",
		"x-rapidapi-host": "tides.p.rapidapi.com",
		"x-rapidapi-key":  tideKey,
		"useQueryString":  "true",
	}

	data := Get(url, "tides", headers)

	if err := json.Unmarshal(data, &structuredData); err != nil {
		log.Printf("JSON Error: %v\n %v", err, data)
	}

	return structuredData
}
