package main

import "fmt"

func mainMenu() {
	fmt.Println("Select Menu:\n1. Tambah Data\n2. Edit Data\n3. Hapus Data\n4. Tampilkan Data")
}

func dataMenu() {
	fmt.Println("Select Options:\n1. Cari Wilayah\n2. Urutkan Data")
}

func searchMenu() {
	fmt.Println("Search Options:\n1. Sequential Search\n2. Binary Search")
}

func sortMenu() {
	fmt.Println("Sort Options:\n1. Selection Sort\n2. Insertion Sort")
}
