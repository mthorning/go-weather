package data

// GeoData from ipGeo API
type GeoData struct {
	Location struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Date                 string  `json:"date"`
	CurrentTime          string  `json:"current_time"`
	Sunrise              string  `json:"sunrise"`
	Sunset               string  `json:"sunset"`
	SunStatus            string  `json:"sun_status"`
	SolarNoon            string  `json:"solar_noon"`
	DayLength            string  `json:"day_length"`
	SunAltitude          float64 `json:"sun_altitude"`
	SunDistance          float64 `json:"sun_distance"`
	SunAzimuth           float64 `json:"sun_azimuth"`
	Moonrise             string  `json:"moonrise"`
	Moonset              string  `json:"moonset"`
	MoonStatus           string  `json:"moon_status"`
	MoonAltitude         float64 `json:"moon_altitude"`
	MoonDistance         float64 `json:"moon_distance"`
	MoonAzimuth          float64 `json:"moon_azimuth"`
	MoonParallacticAngle float64 `json:"moon_parallactic_angle"`
}
