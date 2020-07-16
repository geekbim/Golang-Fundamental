package main

import (
	"fmt"
)

func main() {
	luas, keliling := calculate(2, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

func calculate(panjang int, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	return luas, keliling
}
