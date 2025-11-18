package main
import "fmt"
func main() {
	var orang int

	fmt.Print("Masukan jumlah orang : ")
	fmt.Scan(&orang)

	if orang > 0 {
		motor := orang /2 
		if orang % 2 != 0 {
			motor++
		}
		fmt.Println("Jumlah motor : ", motor)
	}	
}