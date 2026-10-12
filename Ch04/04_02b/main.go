package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber = 3; dayNumber {
	case 1:
		result = "Today is Monday"
	case 2:
		result = "Today is Tuesday"
	case 3:
		result = "Today is Wednesday"
	case 4:
		result = "Today is Thursday"
	case 5:
		result = "Today is Friday"
	default:
		result = "Today is the weekend"
	}
	fmt.Println(result)

	x := 0
	switch {
	case x < 0:
		result = "x is less than 0"
		fallthrough
	case x == 0:
		result = "x is equal to 0"
		fallthrough
	default:
		result = "x is greater than 0"
	}
	fmt.Println(result)
}
