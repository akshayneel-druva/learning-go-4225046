package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	dog.Sound = "Bark"
	dog.Speak()

	q := dog.Speak1()
	fmt.Print(q)

	println(dog.Speak1())
}

type Dog struct {
   Breed string
   Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

func (d Dog) Speak1() string {
	return fmt.Sprintf("%v! %v! %v\n", d.Sound, d.Sound, d.Sound)

}
