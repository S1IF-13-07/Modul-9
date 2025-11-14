package main 

import "fmt"

func main() {
	var bilangan int
	var hasil bool
	fmt.Scan(&bilangan)

	if bilangan == bilangan {
		hasil = bilangan%2 == 0 && bilangan < 0
		hasil = bool(hasil)
	}
	fmt.Println(hasil)
}