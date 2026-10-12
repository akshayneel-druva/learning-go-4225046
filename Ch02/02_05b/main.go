package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Println("Rounded sum:", sum)

	fmt.Println("Value of Pi is ", math.Pi)

	circleRadius := 5.0
	CircleArea := math.Pi * math.Pow(circleRadius, 2)
	fmt.Println("Area of circle with radius", circleRadius, "is", CircleArea)

	fmt.Printf("Circle Area: %.2f\n", CircleArea)

}
