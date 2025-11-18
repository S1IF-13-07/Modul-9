package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Print("bilangan: ")
	fmt.Scanln(&bilangan)
	cek := bilangan%2 == 0 && bilangan < 0
	fmt.Println("genap negatif", cek)
	if bilangan%2 != 0 && bilangan > 0 {
		fmt.Print()
	}
}
