package data

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Longtitude for API requests - currently Truro
const Longtitude = -5.0507

// Latitude for API requests - currently Truro
const Latitude = 49.262951

func fetchNewData(url string, filename string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("New Request:", err)
	}

	req.Header.Set("Accept", "application/json")
	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Sent request: ", err)
	}

	apiData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Read response: ", err)
	}
	resp.Body.Close()

	value, exists := os.LookupEnv("MODE")
	if exists && value == "development" {
		ioutil.WriteFile(fmt.Sprintf("data/cache/%v.json", filename), apiData, 0644)
		if err != nil {
			log.Fatal("Read response: ", err)
		}
	}

	return apiData
}

// Get returns either cached data from file in development or calls the API
func Get(url string, filename string) []byte {
	cachedContent, err := ioutil.ReadFile(fmt.Sprintf("%v.json", filename))
	if err != nil {
		return fetchNewData(url, filename)
	}
	return cachedContent
}
