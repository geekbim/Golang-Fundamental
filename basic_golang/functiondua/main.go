package main

import (
	"fmt"
)

func main() {
	result := add(1, 1)
	fmt.Println(result)
}

func add(number, numberTwo int) int {
	return number + numberTwo
}
