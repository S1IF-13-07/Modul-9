package main

import "fmt"

func main() {
	var orang int

	fmt.Print("Masukkan jumlah orang yang akan touring: ")
	fmt.Scan(&orang)

	motor := orang / 2

	if orang%2 != 0 {
		motor++
	}

	fmt.Printf("Jumlah motor yang diperlukan: %d\n", motor)
}
