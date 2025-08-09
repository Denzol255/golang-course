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

func CreateBin(config *config.Config, name *string, binData *bins.RecordData) (*bins.Bin, error) {
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

	if response.StatusCode != 200 {
		return nil, errors.New("ERROR_WHILE_CREATE_BIN")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result bins.JSONBinResponse
	json.Unmarshal(responseBody, &result)

	return bins.NewBin(result.Metadata.Name, result.Metadata.Id, result.Metadata.Private, result.Record), nil
}

func GetBin(config *config.Config, id *string) (*bins.RecordData, error) {
	client := &http.Client{}

	baseUrl, err := url.Parse(config.PrimaryUrl + "/" + *id)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Master-Key", config.Key)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("ERROR_WHILE_GET_BIN")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var result bins.JSONBinResponse
	json.Unmarshal(responseBody, &result)

	return &result.Record, nil
}

func UpdateBin(config *config.Config, id *string, binData *bins.RecordData) (*string, *bins.RecordData, error) {
	body, _ := json.Marshal(map[string]string{
		"text": binData.Text,
	})
	client := &http.Client{}
	baseUrl, err := url.Parse(config.PrimaryUrl + "/" + *id)
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

func DeleteBin(config *config.Config, id *string) (string, error) {
	client := &http.Client{}

	baseUrl, err := url.Parse(config.PrimaryUrl + "/" + *id)
	if err != nil {
		return "", err
	}

	request, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	if err != nil {
		return "", err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-Master-Key", config.Key)
	response, err := client.Do(request)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		return "", errors.New("ERROR_WHILE_DELETE_BIN")
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	var result bins.JSONBinDeleteResponse
	json.Unmarshal(responseBody, &result)
	return result.Message, nil
}
