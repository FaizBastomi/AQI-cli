package utils

import (
	"errors"
	"strings"
)

func LinearSearch(A []AirPolution, keyword string) ([]AirPolution, error) {
	var i int
	var data []AirPolution

	for i = 0; i < len(A); i++ {
		if strings.Contains(strings.ToLower(A[i].Lokasi), strings.ToLower(keyword)) {
			data = append(data, A[i])
		}
	}

	if len(data) == 0 {
		return []AirPolution{}, errors.New("no data found")
	}

	return data, nil
}

func BinarySearch(A []AirPolution, keyword string) ([]AirPolution, error) {
	var left, mid, right int
	var data []AirPolution
	left = 0
	right = len(A) - 1

	InsSortAscByLokasi(&A)
	for left <= right {
		mid = (left + right) / 2

		if strings.ToLower(A[mid].Lokasi) == strings.ToLower(keyword) {
			data = append(data, A[mid])
			break
		} else if strings.ToLower(A[mid].Lokasi) < strings.ToLower(keyword) {
			left = mid + 1
		} else if strings.ToLower(A[mid].Lokasi) > strings.ToLower(keyword) {
			right = mid - 1
		}
	}

	if len(data) == 0 {
		return []AirPolution{}, errors.New("data not found")
	}

	return data, nil
}
