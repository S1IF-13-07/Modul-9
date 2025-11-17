package main
import "fmt"

func main() {
	var a, b int
	var hasil bool
	fmt.Print("Masukkan bilangan bulat positif pertama: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan bilangan bulat positif kedua: ")
	fmt.Scanln(&b)

	if b == b {
		hasil = b % a == 0
		hasil = bool(hasil)
	}
	fmt.Println(hasil)


	if a == a {
		hasil = a % b == 0
		hasil = bool(hasil)
	}
	fmt.Println(hasil)
}