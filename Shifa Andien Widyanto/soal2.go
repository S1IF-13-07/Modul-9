package main
import "fmt"
func main() {
	var a int
	var teks string

	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&a)

	teks = "bukan"
	if a % 2 ==  0 && a < 0 {
		teks = "genap negatif"
	}
	fmt.Println(teks)
}