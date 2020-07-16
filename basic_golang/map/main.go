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
		"ruby": "is beautiful",
		"go":   "is super fast",
	}

	fmt.Println(myMap)
}
