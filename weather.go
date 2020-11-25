package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type PlaceholderAPI struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func fetchNewData() []byte {
	url := "https://jsonplaceholder.typicode.com/posts"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("New Request:", err)
	}

	client := &http.Client{}

	fmt.Println("Fetching from API")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Sent request: ", err)
	}

	defer resp.Body.Close()

	// reading resp.Body drains it
	apiData, err := ioutil.ReadAll(resp.Body)

	// but we can fill it back up again:
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(apiData))

	// meaning I can send it here
	writeResponseToFile(resp.Body)
	if err != nil {
		log.Fatal("Read response: ", err)
	}

	// and here
	return apiData
}

func writeResponseToFile(apiResponse io.ReadCloser) {
	out, err := os.Create("test.json")
	if err != nil {
		log.Fatal("Create Json File:", err)
	}
	io.Copy(out, apiResponse)
}

func main() {
	var structuredData []PlaceholderAPI
	var data []byte

	cachedContent, err := ioutil.ReadFile("test.json")
	if err != nil {
		newContent := fetchNewData()
		data = newContent
	} else {
		data = cachedContent
	}

	// fmt.Printf("%v : %T", data, data)

	if err := json.Unmarshal(data, &structuredData); err != nil {
		log.Printf("JSON Error: %v\n %v", err, data)
	}

	fmt.Println(data)
}
