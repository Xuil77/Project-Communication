package main

import "fmt"

func main() {
	var x int
	x = 1 + 1
	fmt.Println("New Int", x)
	if x == 10 {
		fmt.Println("Has reach Maximum Power")
	} else {
		fmt.Println("Charging Power")
	}
}
