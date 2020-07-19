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

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func main() {
	user := User{}
	user.ID = 1
	user.FirstName = "Abim"
	user.LastName = "Gatya"
	user.Email = "manyuabim9@gmail.com"
	user.IsActive = true

	user2 := User{}
	user2.ID = 2
	user2.FirstName = "Dhanu"
	user2.LastName = "Ejas"
	user2.Email = "dhanuejas@gmail.com"
	user2.IsActive = true

	// fmt.Println(user)
	// fmt.Println(user2)
	// fmt.Println(user.FirstName)
	// fmt.Println(user2.LastName)

	// displayName1 := DisplayName(user)
	// displayName2 := DisplayName(user2)

	// fmt.Println(displayName1)
	// fmt.Println(displayName2)

	users := []User{user, user2}
	group := Group{"Gamer", user, users, true}

	DisplayGroup(group)
}

func DisplayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("Users name : ")

	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func DisplayName(user User) string {
	result := fmt.Sprintf("Name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}
