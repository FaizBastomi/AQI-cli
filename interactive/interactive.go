package interactive

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/FaizBastomi/AQI-cli-based/utils"
)

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func DataMenu() {
	fmt.Println("Select Options:\n1. Cari Wilayah\n2. Urutkan Data")
}

func SearchMenu() {
	fmt.Println("Search Options:\n1. Sequential Search\n2. Binary Search")
}

func SortMenu() {
	fmt.Println("Sort Options:\n1. Selection Sort\n2. Insertion Sort")
}

func TambahData(A *[]utils.AirPolution) {
	var lokasi, sumberPolusi string
	var IdxUdara int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Lokasi: ")
	scanner.Scan()
	lokasi = scanner.Text()

	fmt.Print("Sumber Polusi: ")
	scanner.Scan()
	sumberPolusi = scanner.Text()

	fmt.Print("Index Udara: ")
	fmt.Scanln(&IdxUdara)

	utils.AddData(A, lokasi, sumberPolusi, IdxUdara)
	ClearConsole()
}

func UbahDataMenu(A *[]utils.AirPolution) {
	var i, idxUdaraBaru int
	var choice, lokasiBaru, sumberBaru string
	var item utils.AirPolution
	var dataPage []utils.AirPolution
	var currentPage int = 1

	ClearConsole()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Silahkan pilih data yang ingin diubah:")
	for {
		dataPage = utils.PaginateData(*A, currentPage)
		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d:\n", currentPage)
			for i, item = range dataPage {
				fmt.Printf("%d. Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					i+1, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println("[n] Next page\n[p] Previous page\n[q] Main Menu, atau masukan nomor data")
		fmt.Print("Select: ")

		scanner.Scan()
		choice = scanner.Text()
		if choice == "q" {
			ClearConsole()
			return
		} else if choice == "n" {
			currentPage++
		} else if choice == "p" {
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		} else {
			index, err := strconv.Atoi(choice)
			if err != nil || index < 1 || index > len(dataPage) {
				fmt.Println("Pilihan tidak valid.")
			} else {
				ClearConsole()
				item = dataPage[index-1]
				fmt.Println("Ubah Data untuk:")
				fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\n", item.Lokasi, item.SumberPolusi, item.IdxUdara)

				fmt.Print("Lokasi (tekan enter untuk skip): ")
				scanner.Scan()
				lokasiBaru = scanner.Text()
				if lokasiBaru != "" {
					item.Lokasi = lokasiBaru
				}

				fmt.Print("Sumber polusi (tekan enter untuk skip): ")
				scanner.Scan()
				sumberBaru = scanner.Text()
				if sumberBaru != "" {
					item.SumberPolusi = sumberBaru
				}

				fmt.Print("Index udara (tekan enter untuk skip): ")
				scanner.Scan()
				choice = scanner.Text()
				if choice != "" {
					idxUdaraBaru, err = strconv.Atoi(choice)
					if err != nil {
						fmt.Println("Index udara tidak valid.")
					} else {
						item.IdxUdara = idxUdaraBaru
						if idxUdaraBaru >= 0 && idxUdaraBaru <= 50 {
							item.TingkatBahaya = "Baik"
						} else if idxUdaraBaru >= 51 && idxUdaraBaru <= 100 {
							item.TingkatBahaya = "Sedang"
						} else if idxUdaraBaru >= 101 && idxUdaraBaru <= 150 {
							item.TingkatBahaya = "Tidak Sehat"
						} else {
							item.TingkatBahaya = "Berbahaya"
						}
					}
				}
				utils.EditData(A, item.Lokasi, item.SumberPolusi, item.IdxUdara, item.AqiID)
			}
		}
		ClearConsole()
	}
}

func HapusData(A *[]utils.AirPolution) {}

func ShowData(A *[]utils.AirPolution) {
	var currentPage, totalPages int
	var item utils.AirPolution
	var dataPage []utils.AirPolution

	ClearConsole()
	scanner := bufio.NewScanner(os.Stdin)
	currentPage = 1
	totalPages = len(*A) / 5
	if len(*A)%5 != 0 {
		totalPages++
	}

	for {
		dataPage = utils.PaginateData(*A, currentPage)
		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for _, item = range dataPage {
				fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					item.Lokasi, item.SumberPolusi, item.IdxUdara, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
				fmt.Println(strings.Repeat("-", 50))
			}
		}
		fmt.Println("[n] Next page\n[p] Previous page\n[q] Main Menu, atau masukan nomor halaman")
		fmt.Print("Select: ")
		scanner.Scan()
		choice := scanner.Text()
		if choice == "q" {
			ClearConsole()
			return
		} else if choice == "n" {
			currentPage++
		} else if choice == "p" {
			if currentPage > 1 {
				currentPage--
			} else {
				fmt.Println("Sudah di halaman pertama.")
			}
		} else {
			page, err := strconv.Atoi(choice)
			if err != nil || page < 1 {
				fmt.Println("Halaman tidak valid.")
			} else {
				currentPage = page
			}
		}
		ClearConsole()
	}
}

func CariData(A *[]utils.AirPolution) {
	var lokasi string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Masukkan nama lokasi yang dicari: ")
	scanner.Scan()
	lokasi = scanner.Text()

	result := utils.SequentialSearch(*A, lokasi)
	if result != nil {
		fmt.Println("Data ditemukan:")
		fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
			result.Lokasi, result.SumberPolusi, result.IdxUdara, result.TingkatBahaya, result.Waktu.Format("02-January-2006 15:04"))
	} else {
		fmt.Println("Data tidak ditemukan.")
	}
	fmt.Println("Tekan Enter untuk kembali...")
	scanner.Scan()
	ClearConsole()
}

func UrutPolusiTerendah(A *[]utils.AirPolution) {
	utils.SortAscendingByIdxUdara(A)
	ShowData(A)
}

func UrutPolusiTertinggi(A *[]utils.AirPolution) {
	utils.SortDescendingByIdxUdara(A)
	ShowData(A)
}