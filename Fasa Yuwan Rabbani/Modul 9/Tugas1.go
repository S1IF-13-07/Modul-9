package main
import "fmt"

func main() {
	var a int
	var b int
	fmt.Print("Masukkan Jumlah Orang :")
	fmt.Scanln(&a)
	
	if a > 0 {
		b = (a + 1) / 2
	}
	
	fmt.Printf("Keluaran: %d motor\n", b)
}