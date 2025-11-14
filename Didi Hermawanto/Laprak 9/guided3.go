package main
import "fmt"
func main() {
	var angka int
	var hasil bool
	fmt.Scan(&angka)
	hasil = (angka%2 == 0) && (angka < 0)
	fmt.Println(hasil)
}