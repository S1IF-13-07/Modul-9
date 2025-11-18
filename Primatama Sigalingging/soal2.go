package main
import "fmt"
func main() {
	var bilangan int
	var teks string
	fmt.Scan(&bilangan)
	teks = "bukan"
	if bilangan < 0 {
		teks = "genap positif"
	}
	fmt.Println(teks)
}