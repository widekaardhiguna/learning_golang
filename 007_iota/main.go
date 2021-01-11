package main

import (
	"fmt"
)

const (
	// Iota is identifier that will increment by 1 everytime its declared
	_ = iota

	// "<<" is bit shifting to left
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	// Print variable
	fmt.Printf("kb : %b\n", kb)
	fmt.Printf("mb : %b\n", mb)
	fmt.Printf("gb : %b\n", gb)
}
