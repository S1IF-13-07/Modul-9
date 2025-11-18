package main

import (
	"fmt"
	"math"
)

func main() {
	var angka float64
	_, err := fmt.Scan(&angka)
	if err != nil {
		return
	}

	hasil := math.Abs(angka)
	fmt.Println(hasil)
}
