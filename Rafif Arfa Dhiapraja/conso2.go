package main

import "fmt"

func main() {
	var bilangan int
	var hasil string
	fmt.Scan(&bilangan)

	hasil = "bukan positif"

	if bilangan > 0 {
		hasil = "positif"
	}

	fmt.Println(hasil)
}