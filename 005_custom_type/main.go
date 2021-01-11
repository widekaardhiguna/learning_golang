package main

import (
	"fmt"
)

func main() {
	var x int

	// Declare a type
	type hotdog int
	var y hotdog

	x = 10
	y = 20

	// Print variable with its format
	fmt.Printf("Type  : %T\nValue : %v\n", x, x)
	fmt.Printf("Type  : %T\nValue : %v\n", y, y)
}
