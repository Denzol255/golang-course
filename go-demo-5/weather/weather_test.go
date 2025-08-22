package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "London"
	geoData := geo.GeoData{
		City: expected,
	}
	format := 3
	got, err := weather.GetWeather(geoData, format)
	if err != nil {
		t.Errorf("Unexpected error %v", err)
	}
	if !strings.Contains(got, expected) {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{
		name:   "Bigger format",
		format: 5,
	},
	{
		name:   "Smaller format",
		format: -1,
	},
	{
		name:   "Zero format",
		format: 0,
	},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			geoData := geo.GeoData{
				City: "London",
			}
			expected := weather.ErrFormatError
			_, err := weather.GetWeather(geoData, tc.format)
			if err != expected {
				t.Errorf("Expected %s, got %s", expected, err)
			}
		})
	}
}
