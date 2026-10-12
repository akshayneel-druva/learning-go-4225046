package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filename := "./test.txt"
	file, err := os.Create(filename)
	checkError(err)
	defer file.Close()

	length, err := io.WriteString(file, "Test 1!")
	fmt.Printf("Wrote file with %v characters", length)

	readFile(filename)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}


func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	fmt.Println("\ndata from file : ",string(data))
}
