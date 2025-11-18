package main

import (
	"fmt"
)

func main() {
	var x int
	if _, err := fmt.Scan(&x); err != nil {
		return
	}

	if x < 0 && x%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}
