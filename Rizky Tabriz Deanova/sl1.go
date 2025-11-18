package main

import "fmt"

func main() {
	var wisatawan int
	fmt.Scan(&wisatawan)
	motor := wisatawan / 2
	if wisatawan%2 != 0 {
		motor += 1
	}
	fmt.Println(motor)
}
