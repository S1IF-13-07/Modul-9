package main
import"fmt"
func main(){
	var x int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	if x > 0 {
	 fmt.Println("Bilangan Positif") 
	} else {
	 fmt.Println("Bukan positif")
	}
} 