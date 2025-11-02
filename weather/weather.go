package weather

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

var ErrInvalidFormat = errors.New("invalid format")

func GetWeather(location geo.LocationData, format int) (string, error) {
	baseUrl, err := url.Parse("https://wttr.in/" + location.City)
	if err != nil {
		fmt.Println("Ошибка при разборе URL:", err)
		return "", err
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	if format > 4 || format < 1 {
		return "", ErrInvalidFormat
	}

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("Ошибка при получении данных о погоде:", err)
		return "", err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return "", err
	}
	defer resp.Body.Close()
	return string(body), nil
}
