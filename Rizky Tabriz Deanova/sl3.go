package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	var cek bool
	if b%a == 0 {
		cek = true
	} else {
		cek = false
	}
	fmt.Println(cek)
	if a%b == 0 {
		cek = true
	} else {
		cek = false
	}
	fmt.Println(cek)
}
