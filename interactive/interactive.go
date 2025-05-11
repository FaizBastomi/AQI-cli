package interactive

import (
	"bufio"
	"fmt"
	"github.com/FaizBastomi/AQI-cli-based/utils"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func MainMenu() {
	fmt.Println("Select Menu:\n1. Tambah Data\n2. Edit Data\n3. Hapus Data\n4. Tampilkan Data\n5. Exit")
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
	var idxPolusi int

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Lokasi: ")
	scanner.Scan()
	lokasi = scanner.Text()

	fmt.Print("Sumber Polusi: ")
	scanner.Scan()
	sumberPolusi = scanner.Text()

	fmt.Print("Index Polusi: ")
	fmt.Scanln(&idxPolusi)

	utils.AddData(A, lokasi, sumberPolusi, idxPolusi)
	ClearConsole()
}

func ShowData(A *[]utils.AirPolution) {
	var currentPage, totalPages int

	scanner := bufio.NewScanner(os.Stdin)
	currentPage = 1
	totalPages = len(*A) / 5
	if len(*A)%5 != 0 {
		totalPages++
	}

	for {
		dataPage := utils.PaginateData(*A, currentPage)
		if len(dataPage) == 0 {
			fmt.Println("Tidak ada data untuk ditampilkan.")
		} else {
			fmt.Printf("Data halaman %d dari %d:\n", currentPage, totalPages)
			for _, item := range dataPage {
				fmt.Printf("Lokasi: %s\nSumber: %s\nIndex: %d\nTingkat: %s\nWaktu: %v\n",
					item.Lokasi, item.SumberPolusi, item.IdxPolusi, item.TingkatBahaya, item.Waktu.Format("02-January-2006 15:04"))
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
