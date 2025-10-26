package geo

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationData struct {
	City   string
	Region string
}

func GetMyLocation(city string) (*LocationData, error) {
	if city != "" {
		return &LocationData{City: city, Region: "SomeRegion"}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("failed to get location data")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	defer resp.Body.Close()
	var locationData LocationData
	err = json.Unmarshal(body, &locationData)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return &locationData, nil
}
