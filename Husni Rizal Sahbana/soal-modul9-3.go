package main

import "fmt"

func main() {
	var x, y int
	var result_1, result_2 bool

	fmt.Scan(&x, &y)

	result_1 = false
	if y%x == 0 {
		result_1 = true
	}

	result_2 = false
	if x%y == 0 {
		result_2 = true
	}

	fmt.Println(result_1)
	fmt.Println(result_2)
}
