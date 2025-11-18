package main
import "fmt"

func main() {
	var x, y int
	var t bool
	fmt.Scan(&x, &y)
	
	if y == y {
		t = y % x == 0
		t = bool(t)
	}
	fmt.Println(t)

	if x == x {
		t = x % y == 0
	    t = bool(t)
	}
	fmt.Println(t)
}