package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"

	fmt.Println("Colors:", colors)
	fmt.Println(colors[2])


	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers:", numbers)
	fmt.Println("length of numbers:", len(numbers))
}
