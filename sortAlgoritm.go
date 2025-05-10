package main

func selectionSort(A *[]airPolution) {
	var i, j, minIdx int
	var temp airPolution

	for i = 0; i < len(*A)-1; i++ {
		minIdx = i
		for j = i + 1; j < len(*A); j++ {
			if (*A)[j].AqiID < (*A)[minIdx].AqiID {
				minIdx = j
			}
		}
		temp = (*A)[i]
		(*A)[i] = (*A)[minIdx]
		(*A)[minIdx] = temp
	}
}

func insertionSort(A *[]airPolution) {
	var i, j int
	var key airPolution

	for i = 1; i < len(*A); i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && (*A)[j].AqiID > key.AqiID {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}
