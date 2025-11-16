package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	genap := n % 2
	hasil := "bukan"
	if genap == 0 && n < 0 {
		hasil = "genap negatif"
	}
	fmt.Print(hasil)
}