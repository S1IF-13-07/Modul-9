package main
import "fmt"

func main() {
	var a int
	fmt.Print("MASUKAN BULANGAN BULATNYA : ")
	fmt.Scanln(&a)

	if a < 0 && a % 2 == 0 {
		fmt.Print("GENAP NEGATIF")
	} else {
		fmt.Print("BUKAN")
	}
}