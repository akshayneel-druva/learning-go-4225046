package main

import (
	"time"
)

func main() {

	go say("Hello from goroutine")
	println("Hello from main function")

	go func(msg string) {
		println(msg)
	}("Hello from anonymous function")


	time.Sleep(2 * time.Second)
	println("Main function is done")
}


func say(msg string) {
	time.Sleep(1 * time.Second)
	println(msg)
}