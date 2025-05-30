package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/FaizBastomi/AQI-cli-based/utils"
)

func subUbahData(A *utils.AirPolutions, B utils.AirPolution) {
	var idxUdaraBaru int
	var lokasiBaru, sumberBaru, choice string
	var scanner *bufio.Scanner
	var err error

	scanner = bufio.NewScanner(os.Stdin)

	fmt.Println("Ubah Data untuk:")
	fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\n", B.Lokasi, B.SumberPolusi, B.IdxUdara)

	fmt.Print("Lokasi (tekan enter untuk skip): ")
	scanner.Scan()
	lokasiBaru = scanner.Text()
	if lokasiBaru != "" {
		B.Lokasi = lokasiBaru
	}

	fmt.Print("Sumber polusi (tekan enter untuk skip): ")
	scanner.Scan()
	sumberBaru = scanner.Text()
	if sumberBaru != "" {
		B.SumberPolusi = sumberBaru
	}

	fmt.Print("Index udara (tekan enter untuk skip): ")
	scanner.Scan()
	choice = scanner.Text()
	if choice != "" {
		idxUdaraBaru, err = strconv.Atoi(choice)
		if err != nil {
			fmt.Println("Index udara tidak valid.")
		} else {
			B.IdxUdara = idxUdaraBaru
			if idxUdaraBaru >= 0 && idxUdaraBaru <= 50 {
				B.TingkatBahaya = "Baik"
			} else if idxUdaraBaru >= 51 && idxUdaraBaru <= 100 {
				B.TingkatBahaya = "Sedang"
			} else if idxUdaraBaru >= 101 && idxUdaraBaru <= 150 {
				B.TingkatBahaya = "Tidak Sehat"
			} else {
				B.TingkatBahaya = "Berbahaya"
			}
		}
	}
	utils.EditData(A, B.Lokasi, B.SumberPolusi, B.IdxUdara, B.AqiID)
	ClearConsole()
}

func subUrutData(A *utils.AirPolutions, sortType int) {
	switch sortType {
	case 0:
		utils.SelSortDescByIdxUdara(A)
	case 1:
		utils.InsSortDescByTime(A)
	}
	ClearConsole()
}
