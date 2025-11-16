package main

import  "fmt"

func main() {
	var b int
	fmt.Scan(&b)
	teks := "Positif"
	if b <= 0 {
		teks = "Bukan Positif"
	}
	fmt.Print(teks)
}