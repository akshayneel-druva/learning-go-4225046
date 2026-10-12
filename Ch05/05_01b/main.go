package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Functions")

	test1()
	added := addValues(3, 4)
	fmt.Printf("The sum of 3 and 4 is %v\n", added)


	addedAll := addAllValues(1, 2, 3, 4, 5)
	fmt.Printf("The sum of all values is %v\n", addedAll)
	fmt.Println("sum is : \n", addedAll)

	sum, count, average := addAllValues1(1, 2, 3, 4, 5)
	fmt.Printf("The sum of all values is %v\n", sum)
	fmt.Printf("The count of all values is %v\n", count)
	fmt.Printf("The average of all values is %v\n", average)

}

func test1() {
	// fmt.Println("This is test1")
	println("This is test1")
}

func addValues(x int, y int) int {
	return x + y
}

func addValues1(x, y int) int {
	return x + y
}

func addAllValues(values ...int) int {
	sum := 0
	for _, value :=range values {
		sum += value
	}
	return sum
}

func addAllValues1(values ...int) (int, int, float64)	{
	sum := 0
	for _, value := range values {
		sum += value
	}
	average := float64(sum) / float64(len(values))
	return sum, len(values), average
}