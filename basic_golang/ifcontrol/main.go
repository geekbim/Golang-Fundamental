package main

import "fmt"

func main() {
	age := 20

	// if
	if age > 10 {
		fmt.Println("Boleh Bermain Game")
	}

	// if else
	if age < 10 {
		fmt.Println("Boleh Bermain Game")
	} else {
		fmt.Println("Kamu Belum Boleh")
	}

	// else if
	score := 90
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score < 60 {
		grade = "D"
	} else if score < 70 {
		grade = "C"
	} else if score < 80 {
		grade = "B"
	} else if score < 90 {
		grade = "B+"
	} else {
		grade = "A"
	}

	fmt.Println("Grade saya adalah ", grade)
}
