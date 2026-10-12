package main

import (
	"fmt"
)

type Person struct {
		Name string
		Age  int
	}

func main() {
	fmt.Println("Structs")

	aPerson := Person{"John", 30}

	fmt.Println("Person:", aPerson)
	fmt.Println("Name:", aPerson.Name)

	aPerson.Age = 31
	fmt.Println("Updated Age:", aPerson.Age)

}
