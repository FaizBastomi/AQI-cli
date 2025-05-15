package utils

import (
	"fmt"
	"time"
)

type AirPolution struct {
	AqiID         string    `json:"aqiID"`
	Lokasi        string    `json:"lokasi"`
	SumberPolusi  string    `json:"sumberPolusi"`
	TingkatBahaya string    `json:"tingkatBahaya"`
	IdxUdara      int       `json:"IdxUdara"`
	Waktu         time.Time `json:"waktu"`
}

func AddData(data *[]AirPolution, lokasi, sumberPolusi string, IdxUdara int) {
	var tingkat string

	if IdxUdara >= 0 && IdxUdara <= 50 {
		tingkat = "Baik"
	} else if IdxUdara >= 51 && IdxUdara <= 100 {
		tingkat = "Sedang"
	} else if IdxUdara >= 101 && IdxUdara <= 150 {
		tingkat = "Tidak Sehat"
	} else {
		tingkat = "Berbahaya"
	}

	*data = append(*data, AirPolution{
		AqiID:  fmt.Sprintf("AQI%d", len(*data)+1),
		Lokasi: lokasi, SumberPolusi: sumberPolusi, IdxUdara: IdxUdara,
		Waktu:         time.Now(),
		TingkatBahaya: tingkat,
	})
}

func EditData(data *[]AirPolution, lokasi, sumberPolusi string, IdxUdara int, aqiID string) {
	var i int
	var user AirPolution

	for i, user = range *data {
		if user.AqiID == aqiID {
			(*data)[i].Lokasi = lokasi
			(*data)[i].SumberPolusi = sumberPolusi
			(*data)[i].IdxUdara = IdxUdara
			(*data)[i].Waktu = time.Now()

			if IdxUdara >= 0 && IdxUdara <= 50 {
				(*data)[i].TingkatBahaya = "Baik"
			} else if IdxUdara >= 51 && IdxUdara <= 100 {
				(*data)[i].TingkatBahaya = "Sedang"
			} else if IdxUdara >= 101 && IdxUdara <= 150 {
				(*data)[i].TingkatBahaya = "Tidak Sehat"
			} else {
				(*data)[i].TingkatBahaya = "Berbahaya"
			}
			break
		}
	}
}

func DeleteData(data *[]AirPolution, aqiID string) {
	var i int
	var user AirPolution

	for i, user = range *data {
		if user.AqiID == aqiID {
			*data = append((*data)[:i], (*data)[i+1:]...)
		}
	}
}
