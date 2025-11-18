package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	if x % 2 != 0 {
		x = (x/2) + 1
	}
	if x % 2 == 0 {
		x = (x/2)
	}
	fmt.Print(x)
}