package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan bilangan X: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan Y: ")
	fmt.Scan(&y)

	fmt.Println(y%x == 0)

	fmt.Println(x%y == 0)
}