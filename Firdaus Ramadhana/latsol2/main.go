package main

import "fmt"

func main() {
	var a int 
	var bilangan string
	fmt.Print("Masukkan nilai bilangan: ")
	fmt.Scan(&a)
	if a%2 == 0 && a < 0 {
		bilangan = "genap negatiif"
	} else {
		bilangan = "bukan"
	}
	fmt.Println(bilangan)
}