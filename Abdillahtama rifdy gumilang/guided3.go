package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	// Cek apakah bilangan genap dan negatif
	if bilangan < 0 && bilangan%2 == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
