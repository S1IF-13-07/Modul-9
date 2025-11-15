package main  
import "fmt"  
func main() {  
	var bilangan int  
	var text string 
	fmt.Scan(&bilangan)  
text = "bukan-positif" 
if bilangan > 0 { 
text = "positif" 
}  
fmt.Println(text)  
}  