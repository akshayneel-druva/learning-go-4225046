package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	str, _ := reader.ReadString('\n')
	fmt.Println("You entered:", str)

	fmt.Print("Enter Number: ")
	str2, _ := reader.ReadString('\n')
	f, err := strconv.ParseFloat(strings.TrimSpace(str2), 64)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("You entered number:", f)
	}


}
