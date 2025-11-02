package geo_test

import (
	"testing"
	"weather/geo"
)

func TestGetMyLocation(t *testing.T) {
	// Arrange
	city := "London"
	expected := geo.LocationData{
		City: "London",
	}
	// Act
	location, err := geo.GetMyLocation(city)
	// Assert
	if err != nil {
		t.Error("Ошибка при получении местоположения:", err)
	}
	if location.City != expected.City {
		t.Errorf("Expected city %s, got %s", expected.City, location.City)
	}
}

func TestGetMyLocation_InvalidCity(t *testing.T) {
	// Arrange
	city := "InvalidCityName"
	// Act
	_, err := geo.GetMyLocation(city)
	// Assert
	if err != geo.ErrInvalidCity {
		t.Errorf("Expected error %v, got %v", geo.ErrInvalidCity, err)
	}
}
