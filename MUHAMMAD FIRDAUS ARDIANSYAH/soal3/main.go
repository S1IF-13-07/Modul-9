package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("masukkan x:")
	fmt.Scan(&x)
	fmt.Print("masukkan y:")
	fmt.Scan(&y)

	if y%x == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if x%y == 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}