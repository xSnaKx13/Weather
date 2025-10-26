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
