package weather

type WeatherResponse struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Humidity  int     `json:"humidity"`
		WindKph   float64 `json:"wind_kph"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
}

type WeatherData struct {
	City        string
	Temperature float64
	WindSpeed   float64
	Humidity    int
}
