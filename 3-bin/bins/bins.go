package bins

import "time"

type Bin struct {
	Id        string
	Private   bool
	CreatedAt time.Time
	Name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin() Bin {
	return Bin{
		Id:        "",
		Private:   false,
		CreatedAt: time.Now(),
		Name:      "",
	}
}

func NewBinList() BinList {
	return BinList{
		Bins: []Bin{},
	}
}
