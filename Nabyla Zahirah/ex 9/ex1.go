package main
import "fmt"
func main(){
	var a int
	fmt.Print("Masukan a: ")
	fmt.Scan(&a)

	if a < 0 {
		a = -a
	}
	fmt.Print(a)
}