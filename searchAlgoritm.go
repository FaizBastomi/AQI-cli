package main

func linearSearch(A []airPolution, aqiID string) int {
	var i int
	var idx int
	idx = -1

	for i = 0; i < len(A); i++ {
		if A[i].AqiID == aqiID {
			idx = i
			break
		}
	}

	return idx
}

func binarySearch(A []airPolution, aqiID string) int {
	var left, mid, right, idx int
	idx = -1
	left = 0
	right = len(A) - 1

	for left <= right && idx == -1 {
		mid = (left + right) / 2

		if A[mid].AqiID == aqiID {
			idx = mid
		} else if A[mid].AqiID < aqiID {
			left = mid + 1
		} else if A[mid].AqiID > aqiID {
			right = mid - 1
		}
	}

	return idx
}
