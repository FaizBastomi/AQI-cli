package main

import (
	"encoding/json"
	"os"
)

func readFromJSON() ([]airPolution, error) {
	dataByte, err := os.ReadFile("data.json")
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(dataByte, &data)

	return data, nil
}

func writeToJSON(data []airPolution) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile("data.json", dataByte, 0644)
	if err != nil {
		return err
	}

	return nil
}
