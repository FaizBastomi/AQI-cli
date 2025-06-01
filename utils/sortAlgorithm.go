package utils

func SelSortDescByIdxUdara(A *AirPolutions) {
	var i, idx, pass int

	for pass = 0; pass < len(*A)-1; pass++ {
		idx = pass
		for i = pass + 1; i < len(*A); i++ {
			if (*A)[i].IdxUdara > (*A)[idx].IdxUdara {
				idx = i
			}
		}
		(*A)[pass], (*A)[idx] = (*A)[idx], (*A)[pass]
	}
}

func InsSortDescByTime(A *AirPolutions) {
	var i, j int
	var key AirPolution

	for i = 1; i < len(*A); i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && (*A)[j].Waktu.Before(key.Waktu) {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}

func InsSortAscByLokasi(A *[]AirPolution) {
	var i, j int
	var key AirPolution

	for i = 1; i < len(*A); i++ {
		key = (*A)[i]
		j = i - 1

		for j >= 0 && (*A)[j].Lokasi > key.Lokasi {
			(*A)[j+1] = (*A)[j]
			j--
		}
		(*A)[j+1] = key
	}
}
