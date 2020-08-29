package main

import (
	"fmt"
)

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// *numberB = 10
	// fmt.Println(numberA)
	// fmt.Println(*numberB)

	var numberA int = 5
	var numberB *int = &numberA

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	numberA = 20
	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)
}
