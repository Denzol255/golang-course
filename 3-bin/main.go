package main

import (
	"app/bin/bins"
	"app/bin/storage"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	bin := bins.NewBin("test2", "2", false)
	binList, err := storage.GetBins("storage.json")
	if err != nil {
		color.Red(err.Error())
		return
	}
	binList.AddBin(*bin)
	storage.SaveBins(binList)
	fmt.Println(binList)
}
