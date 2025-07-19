package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const baseURL = "http://api.weatherapi.com/v1/current.json"

type WeatherService struct {
	apiKey string
}

func NewWeatherService(apiKey string) *WeatherService {
	return &WeatherService{
		apiKey: apiKey,
	}
}

func (s *WeatherService) GetWeather(city string) (*WeatherData, error) {
	url := fmt.Sprintf("%s?key=%s&q=%s&aqi=no", baseURL, s.apiKey, city)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("ошибка при запросе к API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API вернул ошибку %d: %s", resp.StatusCode, string(body))
	}

	var weatherResp WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherResp); err != nil {
		return nil, fmt.Errorf("ошибка при разборе ответа: %w", err)
	}

	weather := &WeatherData{
		City:        weatherResp.Location.Name,
		Temperature: weatherResp.Current.TempC,
		Humidity:    weatherResp.Current.Humidity,
		WindSpeed:   weatherResp.Current.WindKph / 3.6,
	}

	return weather, nil
}
