package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)
	if bilangan > 0 {
		fmt.Println("Bilangan tersebut adalah bilangan positif.")
	} else if bilangan < 0 {
		fmt.Println("Bilangan tersebut adalah bilangan negatif.")
	} else {
		fmt.Println("Bilangan tersebut adalah nol.")
	}
	}