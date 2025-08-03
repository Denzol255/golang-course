package bins

import (
	"app/bin/file"
	"encoding/json"
	"errors"
	"time"
)

type RecordData struct {
	Text string `json:"text"`
}

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `json:"name"`
	RecordData
}

type JSONBinResponse struct {
	Record   RecordData
	Metadata Bin
}

func NewBin(name, id string, private bool, recordData RecordData) *Bin {
	return &Bin{
		Id:         id,
		Private:    private,
		CreatedAt:  time.Now(),
		Name:       name,
		RecordData: recordData,
	}
}

func GetBinDataFromFile(fileName *string) (*RecordData, error) {
	if !file.CheckForJSON(*fileName) {
		return nil, errors.New("FILE_SHOULD_BE_JSON")
	}

	data, err := file.ReadFile(*fileName)
	if err != nil {
		return nil, err
	}

	var binData RecordData
	err = json.Unmarshal(data, &binData)
	if err != nil {
		return nil, err
	}
	return &binData, err
}
