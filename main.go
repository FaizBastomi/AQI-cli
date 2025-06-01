package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/FaizBastomi/AQI-cli-based/interactive"
	"github.com/FaizBastomi/AQI-cli-based/utils"
)

const (
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Reset  = "\033[0m"
)

func exitText() {
	fmt.Printf("%s%s\n", Green, strings.Repeat("-", 50))
	fmt.Printf("%8sTerima kasih telah menggunakan aplikasi ini!\n", Yellow)
	fmt.Printf("%s%s%s\n", Green, strings.Repeat("-", 50), Reset)
}

func main() {
	var opsi int

	// Read data from JSON file
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	data, err := utils.ReadFromJSON(path + "/data.json")
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	interactive.ClearConsole()
	fmt.Printf("%s%s\n", Green, strings.Repeat("-", 50))
	fmt.Printf("%17sSelamat datang di Aplikasi\n", Yellow)
	fmt.Printf("%9sPengelolaan Data Kualitas Udara\n", " ")
	fmt.Printf("%s%s%s\n\n", Green, strings.Repeat("-", 50), Reset)

	for opsi != 7 {
		fmt.Println("Pilih Menu:")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Edit Data")
		fmt.Println("3. Hapus Data")
		fmt.Println("4. Tampilkan Data")
		fmt.Println("5. Cari Data")
		fmt.Println("6. Histori Data")
		fmt.Println("7. Exit")
		fmt.Print("Masukan opsi: ")
		fmt.Scanln(&opsi)

		switch opsi {
		case 1:
			interactive.TambahData(&data)
		case 2:
			interactive.UbahData(&data)
		case 3:
			interactive.HapusData(&data)
		case 4:
			interactive.ShowData(&data)
		case 5:
			interactive.CariData(&data)
		case 6:
			interactive.PeriodikData(&data)
		}

		interactive.ClearConsole()
	}

	exitText()
	// Write data to JSON file
	_ = utils.WriteToJSON(data, path+"/data.json")
}
