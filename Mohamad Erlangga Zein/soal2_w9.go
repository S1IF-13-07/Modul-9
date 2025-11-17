package main 
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka wok: ")
	fmt.Scan(&n)

	if n < 0 && n % 2 == 0 {
		fmt.Print("Genap negatif")
	} else {
		fmt.Print("Bukan")
	}
}