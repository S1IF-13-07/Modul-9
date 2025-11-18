package main
import "fmt"
func main() {
	var orang int
	fmt.Print("Masukan jumlah orang")
	fmt.Scan(&orang)
	var motor int
	if orang%2 == 0 {
		motor = orang / 2
	} else {
		motor = (orang / 2) + 1
	}
	fmt.Println(motor)
	
}