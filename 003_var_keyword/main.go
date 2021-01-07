package main

import (
	"fmt"
)

var y = "Hello World"

// Variable that DECLARED but not ASSIGNED with a value
// will be automatically assigned a ZERO value.
// For example string = "", int = 0, float = 0, bool = false, pointer = nil
var z bool

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
