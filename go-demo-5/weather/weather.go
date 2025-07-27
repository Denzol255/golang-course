package weather

import (
	"demo/weather/geo"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetWeather(geo geo.GeoData, format int) string {
	baseUrl, err := url.Parse("https://wttr.in/" + geo.City)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}
	params := url.Values{}
	params.Add("format", fmt.Sprint(format))
	baseUrl.RawQuery = params.Encode()
	response, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return ""
	}
	return string(body)
}
