package main

import "fmt"

func main() {

	//Cara 1

	// var a int
	// fmt.Print("Input :")
	// fmt.Scan(&a)
	// Output := false
	// if a < 0 && a%2 == 0 {
	// 	Output = true
	// }
	// fmt.Print(Output)

	//

	//Cara 2
	var bilangan int
	var hasil bool
	fmt.Scan(&bilangan)
	hasil = bilangan%2 == 0 && bilangan < 0
	fmt.Println(hasil)
}
