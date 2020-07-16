package main

import "fmt"

func main() {
	// for i := 0; i < 100; i++ {
	// 	fmt.Println("Saya sedang belajar golang")
	// }

	// i := 1
	// for i <= 100 {
	// 	fmt.Println("Saya sedang belajar golang")
	// 	i++
	// }

	title := "Golang the best language"

	for index, letter := range title {
		fmt.Println("index : ", index, "letter : ", string(letter))
	}
}
