package geo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GeoData struct {
	City string `json:"city"`
}

func GetMyLocation(city string) (*GeoData, error) {
	if city != "" {
		return &GeoData{
			City: city,
		}, nil
	}
	response, err := http.Get("http://ipwho.is/")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("Статус код: " + fmt.Sprint(response.StatusCode))
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var geo GeoData
	json.Unmarshal(body, &geo)
	return &geo, nil
}

type CheckCityResponse struct {
	Error bool `json:"error"`
}

func CheckCity(city string) (isError bool) {
	postBody, _ := json.Marshal(map[string]string{
		"city": city,
	})
	response, err := http.Post("https://countriesnow.space/api/v0.1/countries/population/cities", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return false
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return false
	}
	var result CheckCityResponse
	json.Unmarshal(body, &result)
	return result.Error
}
