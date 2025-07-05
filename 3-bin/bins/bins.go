package bins

import (
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins      []Bin     `json:"bins"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewBin(name, id string, private bool) *Bin {
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: time.Now(),
		Name:      name,
	}
}

func (binList *BinList) AddBin(bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}
