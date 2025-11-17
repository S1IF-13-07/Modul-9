package main
import "fmt"
func main(){
	var a int
	
	fmt.Print("Masukkan a: ")
	fmt.Scan(&a)

	if a > 0{
		fmt.Print("positif")
	}	else {
		fmt.Print("bukan positif")
	}
}