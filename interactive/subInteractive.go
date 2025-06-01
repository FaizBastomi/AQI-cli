package interactive

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

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

func subPeriodikData(A *utils.AirPolutions, time string) {
	var currentPage, totalPages, page, i int
	var item utils.AirPolution
	var dataPage, filteredData, periodicFiltered []utils.AirPolution
	var scanner *bufio.Scanner
	var choice string
	var err error

	ClearConsole()
	scanner = bufio.NewScanner(os.Stdin)
	currentPage = 1

	// Sort based on Air Quality
	utils.SelSortDescByIdxUdara(A)

	for {
		filteredData = utils.FilterNonEmpty(*A)
		periodicFiltered = utils.PeriodicFilter(filteredData, time)
		totalPages = len(periodicFiltered) / 5
		if len(periodicFiltered)%5 != 0 {
			totalPages++
		}
		dataPage = utils.PaginateData(periodicFiltered, currentPage)

		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {

			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for i, item = range dataPage {
				fmt.Printf("%d Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					i+1, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}

		fmt.Println("[n] Halaman berikutnya\n[p] Halaman sebelumnya\n[q] Main Menu, atau masukan nomor halaman")
		fmt.Print("Pilih: ")
		scanner.Scan()
		choice = scanner.Text()

		ClearConsole()
		switch choice {
		case "q":
			return
		case "n":
			if currentPage < totalPages {
				currentPage++
			} else {
				fmt.Println("Sudah di halaman terakhir.")
			}
		case "p":
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		default:
			page, err = strconv.Atoi(choice)
			if err != nil || page < 1 {
				fmt.Println("Halaman tidak valid.")
			} else {
				currentPage = page
			}
		}
	}
}
