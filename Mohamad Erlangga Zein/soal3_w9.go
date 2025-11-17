package main
import "fmt"

func main() {
	var x, y int
	var hasil bool

	fmt.Print("Masukkan X: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)
	
	if y == y {
		hasil = y % x == 0
		hasil = bool(hasil)
	}
	fmt.Println(hasil)

	if x == x {
		hasil = x % y == 0
		hasil = bool(hasil)
	}
	fmt.Println(hasil)
}