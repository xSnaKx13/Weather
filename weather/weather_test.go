package weather_test

import (
	"strings"
	"testing"
	"weather/geo"
	"weather/weather"
)

func TestGetWeather(t *testing.T) {
	// Arrange
	location := geo.LocationData{
		City: "InvalidCity",
	}
	format := 3
	// Act
	weatherInfo, _ := weather.GetWeather(location, format)
	// Assert
	if !strings.Contains(weatherInfo, location.City) {
		t.Errorf("Expected %s, got %s", location.City, weatherInfo)
	}
}

func TestGetWeather_InvalidFormat(t *testing.T) {
	var testCases = []struct {
		name   string
		format int
	}{
		{name: "Big format", format: 312},
		{name: "0 format", format: 0},
		{name: "negative format", format: -1},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			location := geo.LocationData{
				City: "London",
			}
			_, err := weather.GetWeather(location, tc.format)
			if err != weather.ErrInvalidFormat {
				t.Errorf("Expected error %v, got %v", weather.ErrInvalidFormat, err)
			}
		})
	}
}
