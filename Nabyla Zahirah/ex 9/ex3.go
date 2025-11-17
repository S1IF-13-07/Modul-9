package main
import "fmt"
func main(){
	var a int
	var hasil bool
	
	fmt.Print("Masukkan a: ")
	fmt.Scan(&a)

	hasil = a % 2 == 0 && a < 0
	fmt.Println(hasil)
}