package weather

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"weather/geo"
)

func GetWeather(location geo.LocationData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + location.City)
	if err != nil {
		fmt.Println("Ошибка при разборе URL:", err)
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("Ошибка при получении данных о погоде:", err)
		return ""
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Ошибка при чтении тела ответа:", err)
		return ""
	}
	defer resp.Body.Close()
	return string(body)
}
