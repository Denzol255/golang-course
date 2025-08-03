package storage

import (
	"app/bin/bins"
	"app/bin/file"
	"encoding/json"
	"errors"
	"slices"
	"time"

	"github.com/fatih/color"
)

type BinList struct {
	Bins      []bins.Bin `json:"bins"`
	UpdatedAt time.Time  `json:"updated_at"`
}

func (binList *BinList) UpdateBinById(id *string, newData *bins.RecordData) {
	binIdx := slices.IndexFunc(binList.Bins, func(bin bins.Bin) bool {
		return bin.Id == *id
	})
	if binIdx != -1 {
		newBin := bins.NewBin(binList.Bins[binIdx].Name, *id, binList.Bins[binIdx].Private, *newData)
		binList.Bins[binIdx] = *newBin
	}
}

func (binList *BinList) SaveBins(fileName string) {
	binList.UpdatedAt = time.Now()
	data, err := json.Marshal(binList)
	if err != nil {
		color.Red(err.Error())
		return
	}
	file.WriteFile(data, fileName)
}

func GetBinList(fileName string) (*BinList, error) {
	if !file.CheckForJSON(fileName) {
		return nil, errors.New("NOT_JSON_FILE")
	}
	data, err := file.ReadFile(fileName)
	if err != nil {
		return &BinList{
			Bins:      []bins.Bin{},
			UpdatedAt: time.Now(),
		}, nil
	}
	var binList BinList
	if err := json.Unmarshal(data, &binList); err != nil {
		color.Red(err.Error())
	}
	return &binList, nil
}

func (binList *BinList) AddBin(bin *bins.Bin) {
	binList.Bins = append(binList.Bins, *bin)
}
