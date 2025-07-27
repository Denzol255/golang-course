package main

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"flag"
	"fmt"
	"io"
	"strings"
)

func main() {
	city := flag.String("city", "Frankfurt am Main", "User's city")
	format := flag.Int("format", 1, "Format weather")
	flag.Parse()
	isError := geo.CheckCity(*city)
	if isError {
		fmt.Println("City not found")
		return
	}
	geoData, err := geo.GetMyLocation(*city)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	weather := weather.GetWeather(*geoData, *format)
	fmt.Println(weather)
}

func testReader() {
	r := strings.NewReader("Hello World")
	block := make([]byte, 4)
	for {
		_, err := r.Read(block)
		fmt.Printf("%q\n", block)
		if err == io.EOF {
			break
		}
	}
}
