package main

import "fmt"

func main() {
	var a, b int
	var hasil1, hasil2 bool
	fmt.Scan(&a, &b)

	if b % a == 0 {
		hasil1 = true
	}
	fmt.Println(hasil1)

	if a % b == 0 {
		hasil2 = true
	}
	fmt.Println(hasil2)

} 