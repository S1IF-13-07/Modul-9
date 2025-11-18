package main
import"fmt"
func main(){
	var x int
	fmt.Print("Masukan bilangan bulat: ")
	fmt.Scan(&x)
		if x < 0 {
			 x = -x
		}
		fmt.Print("Nilai absolut: ", x)
}