package main
import "fmt"
func main() {
	var number int
	var text string
	fmt.Scan(&number)
	text = "bukan negatif genap"

	if number < 0 && (number%2 == 0) {
		text = "negatif genap"
	}
	fmt.Println(text)
}
