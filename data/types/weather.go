package data

// WeatherData structure from MET API
type WeatherData struct {
	Type     string `json:"type"`
	Features []struct {
		Type     string `json:"type"`
		Geometry struct {
			Type        string    `json:"type"`
			Coordinates []float64 `json:"coordinates"`
		} `json:"geometry"`
		Properties struct {
			RequestPointDistance float64 `json:"requestPointDistance"`
			ModelRunDate         string  `json:"modelRunDate"`
			TimeSeries           []struct {
				Time                      string  `json:"time"`
				ScreenTemperature         int     `json:"screenTemperature"`
				MaxScreenAirTemp          float64 `json:"maxScreenAirTemp,omitempty"`
				MinScreenAirTemp          float64 `json:"minScreenAirTemp,omitempty"`
				ScreenDewPointTemperature float64 `json:"screenDewPointTemperature"`
				FeelsLikeTemperature      float64 `json:"feelsLikeTemperature"`
				WindSpeed10M              float64 `json:"windSpeed10m"`
				WindDirectionFrom10M      int     `json:"windDirectionFrom10m"`
				WindGustSpeed10M          float64 `json:"windGustSpeed10m"`
				Max10MWindGust            float64 `json:"max10mWindGust,omitempty"`
				Visibility                int     `json:"visibility"`
				ScreenRelativeHumidity    float64 `json:"screenRelativeHumidity"`
				Mslp                      int     `json:"mslp"`
				UvIndex                   int     `json:"uvIndex"`
				SignificantWeatherCode    int     `json:"significantWeatherCode"`
				PrecipitationRate         int     `json:"precipitationRate"`
				TotalPrecipAmount         int     `json:"totalPrecipAmount"`
				TotalSnowAmount           int     `json:"totalSnowAmount"`
				ProbOfPrecipitation       int     `json:"probOfPrecipitation"`
			} `json:"timeSeries"`
		} `json:"properties"`
	} `json:"features"`
	Parameters []struct {
		TotalSnowAmount struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"totalSnowAmount"`
		ScreenTemperature struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"screenTemperature"`
		Visibility struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"visibility"`
		WindDirectionFrom10M struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"windDirectionFrom10m"`
		PrecipitationRate struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"precipitationRate"`
		MaxScreenAirTemp struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"maxScreenAirTemp"`
		FeelsLikeTemperature struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"feelsLikeTemperature"`
		ScreenDewPointTemperature struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"screenDewPointTemperature"`
		ScreenRelativeHumidity struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"screenRelativeHumidity"`
		ProbOfPrecipitation struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"probOfPrecipitation"`
		WindSpeed10M struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"windSpeed10m"`
		Max10MWindGust struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"max10mWindGust"`
		SignificantWeatherCode struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"significantWeatherCode"`
		TotalPrecipAmount struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"totalPrecipAmount"`
		MinScreenAirTemp struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"minScreenAirTemp"`
		Mslp struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"mslp"`
		WindGustSpeed10M struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"windGustSpeed10m"`
		UvIndex struct {
			Type        string `json:"type"`
			Description string `json:"description"`
			Unit        struct {
				Label  string `json:"label"`
				Symbol struct {
					Value string `json:"value"`
					Type  string `json:"type"`
				} `json:"symbol"`
			} `json:"unit"`
		} `json:"uvIndex"`
	} `json:"parameters"`
}
