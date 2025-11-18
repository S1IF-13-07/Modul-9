package main

import "fmt"

func main() {
    var orang int
    fmt.Print("Masukkan jumlah orang: ")
    fmt.Scan(&orang)

    if orang % 2 != 0 {   
		orang = orang + 1 
	}
	motor := orang / 2  
    fmt.Println(motor)
}