package main

import "fmt"

func main() {
	var jumlahorang, jumlahmotor int

	fmt.Print("MASUKAN TOTAL JUMLAH ORANG: ")
	fmt.Scan(&jumlahorang)

	if jumlahorang%2 == 0 {
		jumlahmotor = jumlahorang / 2
	} else {
		jumlahmotor = (jumlahorang / 2) + 1
	}
	fmt.Println("JUMLAH MOTOR YANG DIBUTUHKAN: ", jumlahmotor)
}
