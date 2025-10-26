package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "", "Specify the city name")
	format := flag.Int("format", 4, "Specify the weather format (1-4)")
	flag.Parse()

	location, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err)
	}
	weatherInfo := weather.GetWeather(*location, *format)
	fmt.Println(weatherInfo)
}
