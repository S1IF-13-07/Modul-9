package main

import "fmt"

func main() {
	// var a int
	// fmt.Print("Input :")
	// fmt.Scan(&a)
	// if a > 0 {
	// 	fmt.Print("Positif")
	// } else {
	// 	fmt.Print("Bukan Positif")
	// }

	var a int
	fmt.Print("Input :")
	fmt.Scan(&a)
	Output := "Bukan Positif"
	if a > 0 {
		Output = "Positif"
	}
	fmt.Print(Output)
}
