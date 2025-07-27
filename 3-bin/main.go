package main

import (
	"app/bin/api"
	"app/bin/bins"
	"app/bin/config"
	"app/bin/storage"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		color.Red("Не удалось загрузить файл .env")
	}
	bin := bins.NewBin("test2", "2", false)
	binList, err := storage.GetBins("storage.json")
	if err != nil {
		color.Red(err.Error())
		return
	}
	binList.AddBin(*bin)
	storage.SaveBins(binList)
	fmt.Println(binList)
	api.GetData(*config.NewConfig())
}
