package main 

import "fmt"

func main() {
	var a int
	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&a)
	if a < 0 {
		a = -a
	}
	fmt.Println("Nilai absolutenya adalah", a)
}