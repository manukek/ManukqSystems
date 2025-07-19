package config

import (
	"os"

	json5 "github.com/yosuke-furukawa/json5/encoding/json5"
)

type Config struct {
	WeatherApiKey string `json:"weatherApiKey"`
}

func Load(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json5.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
