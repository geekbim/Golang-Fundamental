package main

import "fmt"

func main() {

	// Slice cara pertama
	var gamingConsole []string

	gamingConsole = append(gamingConsole, "Playstation4")
	gamingConsole = append(gamingConsole, "Nintendo Switch")
	gamingConsole = append(gamingConsole, "Xbox One")

	fmt.Println(gamingConsole)
	for _, console := range gamingConsole {
		fmt.Println(console)
	}
}
