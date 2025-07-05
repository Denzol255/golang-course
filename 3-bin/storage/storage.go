package storage

import (
	"app/bin/bins"
	"app/bin/file"
	"encoding/json"
	"errors"
	"time"

	"github.com/fatih/color"
)

func SaveBins(binList *bins.BinList) {
	data, err := json.Marshal(binList)
	if err != nil {
		color.Red(err.Error())
		return
	}
	file.WriteFile(data, "storage.json")
	binList.UpdatedAt = time.Now()
}

func GetBins(fileName string) (*bins.BinList, error) {
	if !file.CheckForJSON(fileName) {
		return nil, errors.New("Файл не является JSON файлом")
	}
	data, err := file.ReadFile("storage.json")
	if err != nil {
		return &bins.BinList{
			Bins:      []bins.Bin{},
			UpdatedAt: time.Now(),
		}, nil
	}
	var binList bins.BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		color.Red(err.Error())
	}
	return &binList, nil
}
