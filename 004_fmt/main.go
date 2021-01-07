package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var x = 265

	// Print variable with its format
	fmt.Printf("Type  : %T\nValue : %v\nBinary: %b\n", x, x, x)

	// Print variable and assign it to another variable
	y := fmt.Sprintln("Value of x :", x)

	// Ignoring error for simplicity.
	io.WriteString(os.Stdout, y)
}
