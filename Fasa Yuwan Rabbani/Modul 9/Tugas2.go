package main
import "fmt"

func main() {
	var a int
	fmt.Print("Masukkin bilangan bulat dong : ")
	fmt.Scanln(&a)

	if a < 0 && a % 2 == 0 {
		fmt.Print("Genap Negatif")
	} else {
		fmt.Print("Bukan")
	}
}