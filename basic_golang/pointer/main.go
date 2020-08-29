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

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	// numberA = 20
	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)

	number := 5
	fmt.Println("Alamat memory : ", &number)
	fmt.Println("Nilai Awal : ", number)

	change(&number, 100)

	fmt.Println("Nilai Akhir : ", number)
	fmt.Println("Alamat Memory : ", &number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat memory : ", *old)
	fmt.Println("Di dalam function : ", old)
}
