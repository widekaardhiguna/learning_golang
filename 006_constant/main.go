package main

import (
	"fmt"
)

func main() {
	const x = 12
	const y = 12.64
	const z = "Hello"

	// Print variable with its format
	fmt.Printf("Type  : %T, Value : %v\n", x, x)
	fmt.Printf("Type  : %T, Value : %v\n", y, y)
	fmt.Printf("Type  : %T, Value : %v\n", z, z)
}
