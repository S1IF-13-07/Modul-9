package main
import "fmt"

func main() {
	var x int
	var t bool

	fmt.Scan(&x)

	if x < 0 && x % 2 == 0 {
		t = true
	}
	fmt.Print(t)
}