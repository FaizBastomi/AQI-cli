package main

import (
	"fmt"
	"github.com/FaizBastomi/AQI-cli-based/interactive"
	"github.com/FaizBastomi/AQI-cli-based/utils"
	"os"
)

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
	for opsi != 5 {
		fmt.Println("Select Menu:\n1. Tambah Data\n2. Edit Data\n3. Hapus Data\n4. Tampilkan Data\n5. Exit")
		fmt.Print("Masukan opsi: ")
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			interactive.TambahData(&data)
		case 2:
			interactive.UbahDataMenu(&data)
		case 3:
		case 4:
			interactive.ShowData(&data)
		}
	}

	// Write data to JSON file
	_ = utils.WriteToJSON(data, path+"/data.json")
}
