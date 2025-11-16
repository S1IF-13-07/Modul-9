package main

import "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	pos := b
	hasil := b % 2 == 0
	teks := true
	if pos >= 0 {
		teks = false
	}
	fmt.Print(hasil && teks)
}