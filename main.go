package main

import (
	"flag"
	"fmt"
	"weather/geo"
	"weather/weather"
)

func main() {
	city := flag.String("city", "", "Город пользователя")
	//region := flag.String("region", "", "Регион пользователя")
	format := flag.Int("format", 3, "Формат погоды: 1 - краткий, 2 - средний, 3 - подробный")
	flag.Parse()

	data, err := geo.GetMyLocation(*city)
	if err != nil {
		panic(err)
	}
	fmt.Printf("City: %s, Region: %s\n", data.City, data.Region)
	weatherInfo := weather.GetWeather(*data, *format)
	fmt.Println(weatherInfo)
}
