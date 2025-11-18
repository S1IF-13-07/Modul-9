package main
import "fmt"
func main() {
	var a int

	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&a)

	if a < 0 {
		a = a * -1
	}
	fmt.Print(a)
}