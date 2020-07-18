package main

import (
	"fmt"
)

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Abim"
	user.LastName = "Gatya"
	user.Email = "manyuabim9@gmail.com"
	user.IsActive = true

	fmt.Println(user)
}
