package main
import "fmt"
func main() {
	var a int
	var hasil bool

	fmt.Print("Masukan bilangan bulat : ")
	fmt.Scan(&a)

	hasil = a % 2 ==  0 && a < 0
	fmt.Println(hasil)
}