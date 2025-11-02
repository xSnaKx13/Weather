package showWeather

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/menu"
	"weather/weather"
)

func ShowWeather() {
	userCity := menu.ShowMenu()
	if userCity == "3" {
		fmt.Println("Выход из программы.")
		return
	}

	city := flag.String("city", userCity, "Specify the city name")
	format := flag.Int("format", 4, "Specify the weather format (1-4)")
	flag.Parse()
	location, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err)
	}
	setWeather, err := weather.GetWeather(*location, *format)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(setWeather)
	fmt.Println("Нажмите Enter, чтобы выйти...")
	fmt.Scanln()
}
