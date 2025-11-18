package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("MASKAN X:")
	fmt.Scan(&x)
	fmt.Print("MASUKAN Y:")
	fmt.Scan(&y)

	if y%x == 0 {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
	if x%y == 0 {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}