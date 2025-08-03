package main

import (
	"app/bin/api"
	"app/bin/config"
	"app/bin/storage"
	"flag"
	"fmt"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		color.Red("Не удалось загрузить файл .env")
	}
	config := config.NewConfig()

	isGet := flag.Bool("get", false, "Get bin")
	isList := flag.Bool("list", false, "Get list of bins")
	isCreate := flag.Bool("create", false, "Create bin")
	isUpdate := flag.Bool("update", false, "Update bin")
	isDelete := flag.Bool("delete", false, "Delete bin")
	fileName := flag.String("file", "sample.json", "File name of bin's data")
	binName := flag.String("name", "test", "Bin name")
	binId := flag.String("id", "", "Bin id")

	fmt.Println(*isGet, *isList, *isUpdate, *isDelete, *binId)

	flag.Parse()

	if *isCreate {
		if len(*fileName) != 0 && len(*binName) != 0 {
			newBin, err := api.CreateBin(config, binName, fileName)
			if err != nil {
				color.Red(err.Error())
				return
			}
			binList, err := storage.GetBinList("storage.json")
			if err != nil {
				color.Red(err.Error())
				return
			}
			binList.AddBin(newBin)
			binList.SaveBins("storage.json")
		}
	}

	if *isUpdate {
		if len(*fileName) != 0 && len(*binId) != 0 {
			id, newData, err := api.UpdateBin(config, binId, fileName)
			if err != nil {
				color.Red(err.Error())
				return
			}
			binList, err := storage.GetBinList("storage.json")
			if err != nil {
				color.Red(err.Error())
				return
			}
			binList.UpdateBinById(id, newData)
			binList.SaveBins("storage.json")
		}
	}
}
