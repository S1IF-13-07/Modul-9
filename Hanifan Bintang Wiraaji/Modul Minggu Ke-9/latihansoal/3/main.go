package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(b % a == 0)
	fmt.Print(a % b == 0)
}