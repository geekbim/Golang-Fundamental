package main

import (
	"fmt"
)

func main() {
	sentece := printMyResult("Saya sedang ")
	fmt.Println(sentece)
}

func printMyResult(sentence string) string {
	resultSentence := sentence + "belajar Golang"
	return resultSentence
}
