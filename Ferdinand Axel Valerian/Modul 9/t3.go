package main

import "fmt"

func main() {
    var x, y int
    fmt.Scan(&x, &y)
    
    faktorXY := y % x == 0
    faktorYX := x % y == 0
    
    fmt.Println(faktorXY)
    fmt.Println(faktorYX)
}