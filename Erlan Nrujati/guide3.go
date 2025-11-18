package main
import"fmt"
func main(){
	var x int
	fmt.Print("Masukan bilangan x: ")
	fmt.Scan(&x)
	if (x % 2)==0 && x < 0{
		fmt.Print(true)
	} else {
		fmt.Print(false)
	}
}