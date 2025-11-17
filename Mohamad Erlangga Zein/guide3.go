package main
import "fmt"

func main() {
	var n int
	var b bool

	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	if n < 0 && n % 2 == 0 {
		b = true
	}
	fmt.Print(b)
}