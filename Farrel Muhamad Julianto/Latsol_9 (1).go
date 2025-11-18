package main

import (
	"fmt"
)

func main() {
	var n int

	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	if n <= 0 {
		fmt.Println(0)
		return
	}

	motor := (n + 1) / 2
	fmt.Println(motor)
}
