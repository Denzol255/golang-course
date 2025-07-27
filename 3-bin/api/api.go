package api

import (
	"app/bin/config"
	"fmt"
)

func GetData(config config.Config) {
	if config.Key == "" {
		panic("KEY is not set")
	}
	fmt.Println(config.Key)
}
