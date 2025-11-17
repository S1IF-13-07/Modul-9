package main 
import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	if n > 0{
		fmt.Print("Positif")
	} else {
		fmt.Print("Bukan Positif")
	}
}