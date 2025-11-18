package main
import "fmt"
func main() {
	var x, y int
	fmt.Print("Masukkan x dan y :")
	fmt.Scan(&x, &y)
	var xFaktorY bool
	var yFaktorX bool
	if y % x == 0 {
		xFaktorY = true
	} else {
		xFaktorY = false
	}
	if x % y == 0 {
		yFaktorX = true
	} else {
		yFaktorX = false
	}
	fmt.Println(xFaktorY)
	fmt.Println(yFaktorX)
}