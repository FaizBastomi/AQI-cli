package utils

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func ReadFromJSON(filename string) ([]AirPolution, error) {
	var data []AirPolution

	dataByte, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			var emptyData = []AirPolution{}
			initialData, errM := json.Marshal(emptyData)
			if errM != nil {
				return nil, errM
			}
			err = os.WriteFile(filename, initialData, 0644)
			if err != nil {
				return nil, err
			}
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

func GetNonEmptyInput(scanner *bufio.Scanner, prompt string) string {
	var input string
	for {
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()
		if strings.TrimSpace(input) != "" {
			break
		}
		fmt.Println("Data tidak boleh kosong.")
	}
	return input
}

func GetIntInput(scanner *bufio.Scanner, prompt string) int {
	var input int
	for {
		fmt.Print(prompt)
		scanner.Scan()
		inputStr := scanner.Text()
		if strings.TrimSpace(inputStr) != "" {
			input, _ = strconv.Atoi(inputStr)
			break
		}
		fmt.Println("Data tidak boleh kosong.")
	}
	return input
}

func SortDescendingByIdxUdara(A *[]AirPolution) {
	sort.Slice(*A, func(i, j int) bool {
		return (*A)[i].IdxUdara > (*A)[j].IdxUdara
	})
}

func SortAscendingByIdxUdara(A *[]AirPolution) {
	sort.Slice(*A, func(i, j int) bool {
		return (*A)[i].IdxUdara < (*A)[j].IdxUdara
	})
}

func SequentialSearch(data []AirPolution, target string) *AirPolution {
	for _, item := range data {
		if strings.EqualFold(item.Lokasi, target) {
			return &item
		}
	}
	return nil
}
