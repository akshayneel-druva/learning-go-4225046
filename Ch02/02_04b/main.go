package main

import "fmt"

func main() {
	a1, a2, a3 := 5, 20, 30
	intSum := a1 + a2 + a3
	fmt.Println("Sum of integers:", intSum)

	f1, f2, f3 := 10.5, 20.1, 30.3
	floatSum := f1 + f2 + f3
	fmt.Println("Sum of floats:", floatSum)

	totalSum := float64(intSum) + floatSum
	fmt.Println("Total sum:", totalSum)

	product := float64(a1) * f3
	fmt.Println("Product of a1 and f1:", product)
}
