package api

import (
	"app/bin/bins"
	"app/bin/config"
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

func CreateBin(config *config.Config, name, fileName *string) (*bins.Bin, error) {
	// TODO: вынести в bins, передавать готовую binData
	binData, err := bins.GetBinDataFromFile(fileName)

	if err != nil {
		return nil, err
	}

	body, _ := json.Marshal(map[string]string{
		"text": binData.Text,
	})
	client := &http.Client{}

	request, err := http.NewRequest("POST", config.PrimaryUrl, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Master-Key", config.Key)
	request.Header.Add("X-Bin-Name", *name)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result bins.JSONBinResponse
	json.Unmarshal(responseBody, &result)

	return bins.NewBin(result.Metadata.Name, result.Metadata.Id, result.Metadata.Private, result.Record), nil
}

func GetBin(config config.Config, id string) {

}

func UpdateBin(config *config.Config, id, fileName *string) (*string, *bins.RecordData, error) {
	// TODO: вынести в bins, передавать готовую binData
	binData, err := bins.GetBinDataFromFile(fileName)
	if err != nil {
		return nil, nil, err
	}

	body, _ := json.Marshal(map[string]string{
		"text": binData.Text,
	})
	client := &http.Client{}
	baseUrl, err := url.Parse(config.PrimaryUrl + *id)
	if err != nil {
		return nil, nil, err
	}

	request, err := http.NewRequest("PUT", baseUrl.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, nil, err
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Master-Key", config.Key)
	response, err := client.Do(request)
	if err != nil {
		return nil, nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, nil, errors.New("ERROR_WHILE_UPDATE_BIN")
	}
	return id, binData, nil
}

func DeleteBin(config config.Config, id string) {

}
