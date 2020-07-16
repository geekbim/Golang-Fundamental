package main

import (
	"fmt"
)

func main() {

	// cara pertama
	// var languages [5]string

	// languages[0] = "Go"
	// languages[1] = "PHP"
	// languages[2] = "Javascript"
	// languages[3] = "Java"
	// languages[4] = "Python"

	// fmt.Println(languages)

	// cara kedua
	// languages := [...]string{"Go", "PHP", "Javascript", "Java", "Python"}
	languages := [...]string{
		"Go",
		"PHP",
		"Javascript",
		"Java",
		"Python",
	}
	length := len(languages)

	fmt.Println(languages)
	fmt.Println(length)
	for index, lang := range languages {
		fmt.Println("Index : ", index, "language : ", lang)
	}
}
