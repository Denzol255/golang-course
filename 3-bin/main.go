package main

import (
	"app/bin/bins"
	"app/bin/storage"
	"fmt"
)

func main() {
	bin := bins.NewBin("test", "1", false)
	binList := storage.GetBins()
	binList.AddBin(*bin)
	storage.SaveBins(binList)
	fmt.Println(binList)
}
