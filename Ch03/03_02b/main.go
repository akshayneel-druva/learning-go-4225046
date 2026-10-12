package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	aInt := 12
	var p *int = &aInt

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}

	value1 := 31.2
	pointer1 := &value1
	*pointer1 = 2.5
	fmt.Println("Pointer :", *pointer1)
	fmt.Println("Value :", value1)


}
