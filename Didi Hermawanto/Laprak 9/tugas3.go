package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan dua bilangan (x y): ")
	fmt.Scan(&x, &y)
	
	faktorXY := (y % x == 0)
	faktorYX := (x % y == 0)

	fmt.Println(faktorXY)
	fmt.Println(faktorYX)
}
