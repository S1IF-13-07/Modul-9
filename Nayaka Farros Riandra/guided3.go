package main

import "fmt"

func main() {
	var a int
	var b bool

	fmt.Scan(&a)

	if a<0 && a%2==0 {
		b = true
	}
	if a>0 && a%2!=0 {
		b = false
	}
	fmt.Println(b)
}