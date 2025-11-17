package main

import "fmt"

func main() {
	var jumlahorang, jumlahmotor int

	fmt.Print("masukkan jumlah orang: ")
	fmt.Scan(&jumlahorang)

	if jumlahorang%2 == 0 {
		jumlahmotor = jumlahorang / 2
	} else {
		jumlahmotor = (jumlahorang / 2) + 1
	}
	fmt.Println("jumlah motor yang dibutuhkan: ", jumlahmotor)
}