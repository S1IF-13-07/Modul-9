package main 
import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	if x > 0{
		fmt.Print("positif")
	} else {
		fmt.Print("bukan positif")
	}
}