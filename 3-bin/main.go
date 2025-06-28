package main

import (
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func main() {
	bin := newBin()
	fmt.Println(bin)
}

func newBin() Bin {
	return Bin{
		id:        "",
		private:   false,
		createdAt: time.Now(),
		name:      "",
	}
}
