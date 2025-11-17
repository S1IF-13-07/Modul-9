package main
import "fmt"

func main(){
	var n int
	fmt.Print("Masukkan n: ")
	fmt.Scan(&n)

	if n < 0{
		n *= -1
	}
	fmt.Print(n)
}