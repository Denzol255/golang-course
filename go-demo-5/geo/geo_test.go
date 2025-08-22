package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	got, err := geo.GetMyLocation(city)

	if err != nil {
		t.Error("Ошибка вызова функции GetMyLocation", err)
	}
	if got.City != expected.City {
		t.Errorf("Expected %s, got %s", expected.City, got.City)
	}
}

func TestGetMyLocationNoCity(t *testing.T) {
	city := "Jack London"
	_, err := geo.GetMyLocation(city)

	if err != geo.ErrCityNotFound {
		t.Errorf("Expected %s, got %s", geo.ErrCityNotFound, err)
	}
}
