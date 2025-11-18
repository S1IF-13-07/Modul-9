package main
import (
	"fmt"
	"math"
)

func main(){
	var orang float64
	fmt.Scan(&orang)

	if orang == orang {
		orang /= 2
		motor := math.Round(orang)
		fmt.Print(motor)
	} else {
		fmt.Print("bukan")
	}
}