package main

import "fmt"

func main() {
	var bilangan int
	var hasil bool
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)
	if bilangan % 2 == 0 {
		if bilangan > 0 {
			hasil = true
			fmt.Print("Output: %t (Genap Positif)\n", hasil)
		} else if bilangan < 0 {
			hasil = false
			fmt.Printf("Output: %t (Grnap Negatif\n", hasil)
		} else {
			fmt.Println("Output: Bilangan adalah Nol (Genap Netral).")
		} else {
			fmt.Println("Output: Bilangan adalah Ganjil (Tidak memenuhi kriteria True/False).")
		}
		}
		}