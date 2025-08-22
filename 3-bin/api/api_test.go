package api_test

import (
	"app/bin/api"
	"app/bin/bins"
	"app/bin/config"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

var sampleBinName = "testBin"
var sampleBinData = "some data"
var updateBinData = "updated data"

func init() {
	fmt.Println("INIT_START")
	godotenv.Load("../.env")
	fmt.Println("INIT_END")
}
func TestCreateBin(t *testing.T) {
	config := config.NewConfig()
	got, createErr := api.CreateBin(config, &sampleBinName, &bins.RecordData{
		Text: sampleBinData,
	})

	if createErr != nil {
		t.Errorf("Unexpected error %v while create", createErr)
		return
	}

	if got.Name != sampleBinName {
		t.Errorf(
			"Expected %s, got %s",
			sampleBinName,
			got.Name,
		)
	}

	_, deleteErr := api.DeleteBin(config, &got.Id)

	if deleteErr != nil {
		t.Errorf("Unexpected error %v while delete", createErr)
		return
	}
}

func TestGetBin(t *testing.T) {
	config := config.NewConfig()
	got, createErr := api.CreateBin(config, &sampleBinName, &bins.RecordData{
		Text: sampleBinData,
	})

	if createErr != nil {
		t.Errorf("Unexpected error %v while create", createErr)
		return
	}

	data, getErr := api.GetBin(config, &got.Id)

	if getErr != nil {
		t.Errorf("Unexpected error %v while get", getErr)
		return
	}

	if data.Text != sampleBinData {
		t.Errorf(
			"Expected %s, got %s",
			sampleBinData,
			data.Text,
		)
	}

	_, deleteErr := api.DeleteBin(config, &got.Id)

	if deleteErr != nil {
		t.Errorf("Unexpected error %v while delete", createErr)
		return
	}
}

func TestUpdateBin(t *testing.T) {
	config := config.NewConfig()
	got, createErr := api.CreateBin(config, &sampleBinName, &bins.RecordData{
		Text: sampleBinData,
	})

	if createErr != nil {
		t.Errorf("Unexpected error %v while create", createErr)
		return
	}

	_, data, updateErr := api.UpdateBin(config, &got.Id, &bins.RecordData{
		Text: updateBinData,
	})

	if updateErr != nil {
		t.Errorf("Unexpected error %v while update", updateErr)
		return
	}

	if data.Text != updateBinData {
		t.Errorf(
			"Expected %s, got %s",
			updateBinData,
			data.Text,
		)
	}

	_, deleteErr := api.DeleteBin(config, &got.Id)

	if deleteErr != nil {
		t.Errorf("Unexpected error %v while delete", createErr)
		return
	}
}

func TestDeleteBin(t *testing.T) {
	config := config.NewConfig()
	got, createErr := api.CreateBin(config, &sampleBinName, &bins.RecordData{
		Text: sampleBinData,
	})

	if createErr != nil {
		t.Errorf("Unexpected error %v while create", createErr)
		return
	}

	_, deleteErr := api.DeleteBin(config, &got.Id)

	if deleteErr != nil {
		t.Errorf("Unexpected error %v while delete", createErr)
		return
	}

	_, getErr := api.GetBin(config, &got.Id)

	if getErr != api.ErrNotFoundBin {
		t.Errorf(
			"Expected %s, got %s",
			api.ErrNotFoundBin,
			getErr,
		)
	}
}
