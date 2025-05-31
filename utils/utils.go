package utils

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func FilterNonEmpty(data AirPolutions) []AirPolution {
	var entry AirPolution
	var n int = 0

	for _, entry = range data {
		if entry.AqiID != "" {
			data[n] = entry
			n++
		}
	}

	return data[:n]
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
	for strings.TrimSpace(input) == "" {
		fmt.Print(prompt)
		scanner.Scan()
		input = scanner.Text()

		if strings.TrimSpace(input) == "" {
			fmt.Println("Data tidak boleh kosong.")
		}
	}
	return input
}

func GetIntInput(scanner *bufio.Scanner, prompt string) int {
	var input int
	var inputTrim string
	var err error

	for strings.TrimSpace(inputTrim) == "" {
		fmt.Print(prompt)
		scanner.Scan()
		inputTrim = scanner.Text()

		if strings.TrimSpace(inputTrim) == "" {
			fmt.Println("Data tidak boleh kosong.")
		}

		if input, err = strconv.Atoi(inputTrim); err != nil {
			fmt.Println("Input tidak valid. Harap masukkan angka.")
			inputTrim = ""
		}
	}
	return input
}

func PeriodicFilter(data []AirPolution, period string) []AirPolution {
	var duration time.Duration
	var now, threshold time.Time
	var entry AirPolution

	now = time.Now()
	switch period {
	case "day":
		duration = 3 * 24 * time.Hour
	case "week":
		duration = 7 * 24 * time.Hour
	case "month":
		duration = 30 * 24 * time.Hour
	default:
		fmt.Println("Invalid period")
		return nil
	}

	threshold = now.Add(-duration)
	var result []AirPolution
	for _, entry = range data {
		// Assuming entry.Waktu is of type time.Time
		if entry.Waktu.After(threshold) {
			result = append(result, entry)
		}
	}
	return result
}
