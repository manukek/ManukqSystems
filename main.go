package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/manukek/manukqsystems/config"
	"github.com/manukek/manukqsystems/weather"
)

func urlFor(name string) string {
	switch name {
	case "static":
		return "/static/"
	default:
		return "/"
	}
}

func main() {
	cfg, err := config.Load("config.json5")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	weatherSvc := weather.NewWeatherService(cfg.WeatherApiKey)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	tmpl := template.Must(template.New("index.html").Funcs(template.FuncMap{"url_for": urlFor}).ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		weatherData, err := weatherSvc.GetWeather("Taldykorgan")
		if err != nil {
			log.Printf("Error getting weather: %v", err)
			http.Error(w, "Ошибка при получении данных о погоде", http.StatusInternalServerError)
			return
		}

		data := struct {
			Weather *weather.WeatherData
		}{
			Weather: weatherData,
		}

		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			log.Printf("Template execution error: %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
