package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationData struct {
	City   string
	Region string
}

type PopulationCityResponse struct {
	Error bool `json:"error"`
}

var ErrInvalidCity = errors.New("invalid city name")
var ErrNot200 = errors.New("NOT200")

func GetMyLocation(city string) (*LocationData, error) {
	if city != "" {
		isCity := checkCity(city)
		if !isCity {
			return nil, ErrInvalidCity
		}
		return &LocationData{City: city, Region: "SomeRegion"}, nil
	}
	resp, err := http.Get("https://ipapi.co/json/")
	if err != nil {
		return nil, errors.New(err.Error())
	}
	if resp.StatusCode != 200 {
		return nil, ErrNot200
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

func checkCity(city string) bool {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	resp, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil || resp.StatusCode != 200 {
		return false
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	var populationResp PopulationCityResponse
	err = json.Unmarshal(body, &populationResp)
	if err != nil {
		return false
	}
	return !populationResp.Error
}
