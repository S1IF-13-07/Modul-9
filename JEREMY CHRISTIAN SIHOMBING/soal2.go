package main
import "fmt"
func main() {

    var bilangan int
    var teks string
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)
    teks = "bukan"
    if bilangan % 2 == 0 && bilangan < 0 {
        teks = "genap negatif"
    }
    fmt.Println("Bilangan adalah genap negatif atau bukan?", teks)
}
