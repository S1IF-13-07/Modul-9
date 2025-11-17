package main 

import "fmt"

func main() {
	var a int 
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)
	if a <= 0 {
		fmt.Println("Bilangan tersebut bukan positif")
	} else {
		fmt.Println("Bilangan tersebut positif")
	}
}