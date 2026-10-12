package main

import (
	"fmt"
)

func main() {
	// This is an array
	var colors = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	// This is a slice
	var numbers = []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("Length of numbers:", len(numbers))

	var fruits = make([]string, 0, 3)
	fruits = append(fruits, "Apple", "Banana", "Cherry")

	fmt.Println("Fruits:", fruits)
	fruits = append(fruits, "Date")
	fmt.Println("Fruits:", fruits)

	fruits = remove(fruits, 1)
	fmt.Println("Fruits after removal:", fruits)
}


func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
