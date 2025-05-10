package main

import (
	"fmt"
	"time"
)

type airPolution struct {
	AqiID         string    `json:"aqiID"`
	Lokasi        string    `json:"lokasi"`
	SumberPolusi  string    `json:"sumberPolusi"`
	TingkatBahaya string    `json:"tingkatBahaya"`
	IdxPolusi     int       `json:"idxPolusi"`
	Waktu         time.Time `json:"waktu"`
}

var data []airPolution

func addData(lokasi, sumberPolusi string, idxPolusi int) {
	var tingkat string

	if idxPolusi >= 0 && idxPolusi <= 50 {
		tingkat = "Baik"
	} else if idxPolusi >= 51 && idxPolusi <= 100 {
		tingkat = "Sedang"
	} else if idxPolusi >= 101 && idxPolusi <= 150 {
		tingkat = "Tidak Sehat"
	} else {
		tingkat = "Berbahaya"
	}

	data = append(data, airPolution{
		AqiID:  fmt.Sprintf("AQI%d", len(data)+1),
		Lokasi: lokasi, SumberPolusi: sumberPolusi, IdxPolusi: idxPolusi,
		Waktu:         time.Now(),
		TingkatBahaya: tingkat,
	})
}

func editData(lokasi, sumberPolusi string, idxPolusi int, aqiID string) {
	var i int
	var user airPolution

	for i, user = range data {
		if user.AqiID == aqiID {
			data[i].Lokasi = lokasi
			data[i].SumberPolusi = sumberPolusi
			data[i].IdxPolusi = idxPolusi
		}
	}
}

func deleteData(aqiID string) {
	var i int
	var user airPolution

	for i, user = range data {
		if user.AqiID == aqiID {
			data = append(data[:i], data[i+1:]...)
		}
	}
}
