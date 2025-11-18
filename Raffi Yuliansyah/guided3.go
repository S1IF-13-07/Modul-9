package main

import "fmt"

func main(){
	var x int
	var b bool
	b = false
	fmt.Scan(&x)
	if x < 0 && x % 2 == 0{
		b = true
	}
	fmt.Print(b)
	
}