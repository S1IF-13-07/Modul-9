package main
import "fmt"
func main() {
	var x, y int

	fmt.Print("Masukan dua bilangn positif : ")
	fmt.Scan(&x, &y)

	fmt.Println(y % x == 0)
	fmt.Println(x % y == 0)
}