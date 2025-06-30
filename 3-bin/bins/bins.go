package bins

import "time"

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func NewBin() Bin {
	return Bin{
		id:        "",
		private:   false,
		createdAt: time.Now(),
		name:      "",
	}
}

func NewBinList() BinList {
	return BinList{
		bins: []Bin{},
	}
}
