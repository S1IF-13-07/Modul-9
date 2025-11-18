package main

import "fmt"

func main() {
	var a int
	fmt.Print("Input :")
	fmt.Scan(&a)
	if a < 0 {
		a = a * -1
		//a = -a
	}
	fmt.Print("Output :", a)
}
