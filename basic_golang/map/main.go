package main

import (
	"fmt"
)

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["Javascript"] = 8
	// myMap["go"] = 10

	// fmt.Println(myMap["ruby"])

	myMap := map[string]string{
		"ruby":       "is beautiful",
		"go":         "is super fast",
		"javascript": "is hype",
	}

	fmt.Println(myMap)
	for key, val := range myMap {
		fmt.Println("Key : ", key, "Value : ", val)
	}

	fmt.Println("=========================")

	delete(myMap, "ruby")
	for key, val := range myMap {
		fmt.Println("Key : ", key, "Value : ", val)
	}

	fmt.Println("=========================")

	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}
