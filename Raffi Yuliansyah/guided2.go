package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)
	if x<=0{
		fmt.Print("Bukan Positif")
	}
	if x>0{
		fmt.Print("Positif")
	}
}