package main 

import "fmt"

func main() {
	var a int 
	var hasil bool
	fmt.Print("Masukkan Bilangan Bulat: ")
	fmt.Scan(&a)
	if a%2 == 0 {
		hasil = true
	} else {
		hasil = false
	}
	fmt.Println(hasil)
}