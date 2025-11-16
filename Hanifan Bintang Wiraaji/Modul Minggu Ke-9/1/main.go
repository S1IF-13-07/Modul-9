package main

import  "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	if b < 0 {
		b = -b
	}
	fmt.Print(b)
}