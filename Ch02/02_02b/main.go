package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	anum := 10
	stringlen, err := fmt.Println("anum:", anum)

	if err == nil {
		fmt.Println("stringlen:", stringlen)
	}

	fmt.Println(str1, str2, str3)

	fmt.Printf("Value of number: %v \n", anum)
	fmt.Printf("Data type of number: %T \n", anum)
}
