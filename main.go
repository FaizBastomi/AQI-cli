package main

import "fmt"

func main() {
	var opsi int

	// Read data from JSON file
	_, err := readFromJSON()
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	mainMenu()
	fmt.Print("Masukan opsi:")
	fmt.Scanf("%d", &opsi)
	for opsi != 5 {
		switch opsi {
		case 1:
			break
		case 2:
			break
		case 3:
			break
		case 4:
			break
		}
	}

	// Write data to JSON file
	_ = writeToJSON(data)
}
