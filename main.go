package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/manukek/ManukqSystems/config"
	"github.com/manukek/ManukqSystems/weather"
)

func urlFor(name string) string {
	switch name {
	case "static":
		return "/static/"
	default:
		return "/"
	}
}

func nowInUTC() time.Time {
	loc := time.FixedZone("UTC+3", 3*60*60) // поменяйте на нужный вам часовой пояс, в моём случае это UTC+5
	// Если вам нужен другой часовой пояс, измените "UTC+5" и смещение в секундах
	// Например, для UTC+3 используйте "UTC+3", 3*60*60, это время для МСК
	return time.Now().In(loc)
}

func main() {
	cfg, err := config.Load("config.json5")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	weatherSvc := weather.NewWeatherService(cfg.WeatherApiKey)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	funcMap := template.FuncMap{
		"url_for": urlFor,
		"now5":    nowInUTC,
	}
	tmpl := template.Must(template.New("index.html").Funcs(funcMap).ParseFiles("templates/index.html"))

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
