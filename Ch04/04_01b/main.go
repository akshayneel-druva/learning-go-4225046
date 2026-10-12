package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	// x := 3
	var result string

	if x:= 42;x < 3 {
		result = "x is less than 3"
	} else if x == 3 {
		result = "x is equal to 3"
	} else {
		result = "x is greater than 3"
	}
	fmt.Println(result)
}
