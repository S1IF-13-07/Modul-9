package main
import "fmt"
func main(){
	var a int

	
	fmt.Print("Masukkan a: ")
	fmt.Scan(&a)

	if a % 2 == 0 && a < 0 {
		fmt.Println("genap negatif")
	}	else {
		fmt.Println("bukan")
	}
}