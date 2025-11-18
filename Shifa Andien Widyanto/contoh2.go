package main
import "fmt"
func main() {
	var a int
	var teks string

	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&a)

	teks = "bukan positif"
	if a > 0 {
		teks = "positif"
	}
	fmt.Println(teks)
}