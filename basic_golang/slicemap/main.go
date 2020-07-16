package main

import "fmt"

func main() {
	students := []map[string]string{
		{"name": "Abim", "score": "A"},
		{"name": "Dhanu", "score": "A"},
		{"name": "Ejas", "score": "A"},
	}

	fmt.Println(students)

	fmt.Println("==========================")

	for _, student := range students {
		fmt.Println(student)
		fmt.Println(student["name"], " - ", student["score"])
	}
}
