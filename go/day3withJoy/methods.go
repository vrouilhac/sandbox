package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	// user := User.Create("Vincent", "crysdev@gmail.com", 23)
	user := User { "Vincent", "cyrs@gmail.com", 23 }

	user.ToString()
	user.ChangeUsername("Julius")
	user.ToString()


	cat := Cat { "Black" }
	duck := Duck { "White" }

	fmt.Println(cat)

	animals := [2]Animal{ cat, duck }

	for _, v := range animals {
		v.Move()
		v.Eat()
	}
}

type User struct {
	username, email string
	age int
}

// func (user User) Create(username, email string, age int) User {
// 	newUser := User { username, email, age }
// 	return newUser
// }

func (user User) ToString() {
	fmt.Println(user)
}

func (user *User) ChangeUsername(username string) {
	user.username = username
}
// REMINDER: You can only declare a method with a receiver whose type is defined in the same package as the method
// Don't know how interface are usefull there ?? ok i do now.. x)

type Animal interface {
	Move()
	Eat() // struct need to implement ALL interface methods to be recognize as implementation
}

type Cat struct {
	color string
}

// Implicit implementation of animal
func (cat Cat) Move() {
	fmt.Println("I'm walking...")
}

func (cat Cat) Eat() {
	fmt.Println("I eat what my slave human do for me because as a cat, i'm the true master xD")
}

func (cat Cat) String() string {
	return fmt.Sprintf("I'm a %v cat!", cat.color)	
}

type Duck struct {
	color string
}

func (duck Duck) Move() {
	fmt.Println("I'm swimming...") // 1 m or 2 ?
}

func (duck Duck) Eat() {
	fmt.Println("I eat something i can't name because the developer who wrote this line didn't know the answer...") // 1 m or 2 ?
}
