package main

import "fmt"

func main() {
	var bilangan int
	var text string
	fmt.Scan(&bilangan)
	text = "bukan"

	if bilangan % 2 == 0 && bilangan < 0 {
		text = "genap negatif"
	}

	fmt.Println(text)
}