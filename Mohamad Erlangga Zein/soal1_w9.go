package main
import (
	"fmt"
	"math"
)

func main(){
	var jumlahOrang float64
	fmt.Print("Masukkan jumlah orang: ")
	fmt.Scan(&jumlahOrang)

	if jumlahOrang == jumlahOrang {
		jumlahOrang /= 2
		motor := math.Round(jumlahOrang)
		fmt.Print("Jumlah motor: ", motor)
	} else {
		fmt.Print("Nggak sesuai bang")
	}
}