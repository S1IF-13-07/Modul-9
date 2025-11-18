package main

import "fmt"

func main(){
	var x, y int
	var n, m bool
	n, m = false, false
	fmt.Scan(&x,&y)
	if y % x == 0 {
		n = true
	}
	if x % y == 0 {
		m = true
	}
	fmt.Println(n)
	fmt.Println(m)
}