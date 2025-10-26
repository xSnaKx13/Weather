package showWeather

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/menu"
	promptdata "weather/promptData"
	"weather/weather"
)

func ShowWeather() {
	var userCity string
	var isWorking bool = true
	for isWorking {
		switch menu.ShowMenu() {
		case "1":
			userCity = ""
			isWorking = false
		case "2":
			userCity = promptdata.PromptData("Введите название города:")
			isWorking = false
		case "3":
			return
		default:
			fmt.Println("\nНекорректный выбор.\nПожалуйста, выберите 1, 2 или 3.")
		}
	}

	city := flag.String("city", userCity, "Specify the city name")
	format := flag.Int("format", 4, "Specify the weather format (1-4)")
	flag.Parse()
	location, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err)
	}
	setWeather := weather.GetWeather(*location, *format)
	fmt.Println(setWeather)
}
