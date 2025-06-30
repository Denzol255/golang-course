package main

import (
	"app/bin/bins"
	"fmt"
)

func main() {
	bin := bins.NewBin()
	binList := bins.NewBinList()
	fmt.Println(bin)
	fmt.Println(binList)
}
