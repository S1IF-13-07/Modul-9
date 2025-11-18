package main  
import "fmt"  
func main() {  
    
    var bilangan int  
    var hasil bool 
    fmt.Print("Masukkan bilangan bulat: ")
    fmt.Scan(&bilangan)  
    hasil = bilangan % 2 == 0 && bilangan < 0 
    fmt.Println("Apakah bilangan genap negatif?", hasil)  
}  