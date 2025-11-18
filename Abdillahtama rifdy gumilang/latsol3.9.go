package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	// Baris pertama: x faktor dari y?
	fmt.Println(y%x == 0)

	// Baris kedua: y faktor dari x?
	fmt.Println(x%y == 0)
}
