package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")

	t := time.Date(2026, time.October, 12, 12, 0, 0, 0, time.Local)
	fmt.Printf("Date: %s\n", t)
	fmt.Printf("Year: %d\n", t.Year())
	fmt.Printf("Type: %T\n", t)

	now := time.Now()
	fmt.Printf(now.Format(time.ANSIC))

	dateFormat := "2006-01-02 15:04:05"
	fmt.Println("\nDate format:", now.Format(dateFormat))

}
