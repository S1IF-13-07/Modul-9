package main

import "fmt"

func main(){
	var x int
	var t string
	fmt.Scan(&x)
	t = "bukan"
	if x < 0 && x % 2 == 0{
		t = "genap negatif"
	}
	fmt.Print(t)
}