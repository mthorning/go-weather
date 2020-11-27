package data

import "time"

// TidesData from the tides API
type TidesData struct {
	Disclaimer string  `json:"disclaimer"`
	Status     int     `json:"status"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Origin     struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Distance  float64 `json:"distance"`
		Unit      string  `json:"unit"`
	} `json:"origin"`
	Datums struct {
		LAT float64 `json:"LAT"`
		HAT float64 `json:"HAT"`
	} `json:"datums"`
	Timestamp int       `json:"timestamp"`
	Datetime  time.Time `json:"datetime"`
	Unit      string    `json:"unit"`
	Timezone  string    `json:"timezone"`
	Datum     string    `json:"datum"`
	Extremes  []struct {
		Timestamp int       `json:"timestamp"`
		Datetime  time.Time `json:"datetime"`
		Height    float64   `json:"height"`
		State     string    `json:"state"`
	} `json:"extremes"`
	Heights []struct {
		Timestamp int       `json:"timestamp"`
		Datetime  time.Time `json:"datetime"`
		Height    float64   `json:"height"`
		State     string    `json:"state"`
	} `json:"heights"`
	Copyright string `json:"copyright"`
}
