package utils

import (
	"encoding/json"
	"os"
)

func ReadFromJSON(filename string) ([]AirPolution, error) {
	var data []AirPolution

	dataByte, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			var emptyData = []AirPolution{}
			initialData, _ := json.Marshal(emptyData)
			_ = os.WriteFile(filename, initialData, 0644)
			return emptyData, nil
		}
		return data, err
	}

	err = json.Unmarshal(dataByte, &data)

	return data, nil
}

func WriteToJSON(data []AirPolution, filename string) error {
	dataByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, dataByte, 0644)
	if err != nil {
		return err
	}

	return nil
}

func PaginateData(data []AirPolution, page int) []AirPolution {
	var perPage, start, end int
	perPage = 5
	start = (page - 1) * perPage
	if start >= len(data) {
		return []AirPolution{}
	}
	end = start + perPage
	if end > len(data) {
		end = len(data)
	}
	return data[start:end]
}
