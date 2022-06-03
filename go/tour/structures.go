package main

import (
	"fmt"
)

type User struct {
	username string
	country string
}

func main() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	vSlice := a[1:4]
	var anotherSlice []int = a[3:7]

	anotherSlice[0] = 20

	fmt.Println(a, vSlice, anotherSlice)

	users := []User{
		{username: "Vincent", country: "FR"},
		{"Julius", "ES"},
	}

	// users[3] = User { // Take care declaring slices doesn't throw error before runtime
	// 	username: "Marie",
	// 	country: "FR",
	// }

	me := users[:1]

	fmt.Println(users, me)

	aLength := len(a)
	var aCap int = cap(a)
	var meCap int = cap(me)
	var meLength int = len(me)

	var n []int

	fmt.Println(aLength, aCap, meCap, meLength, n, len(n), cap(n))
}



type Person struct {
	Name string
	Age int
	Score float64
}

func oldmain() {
	me := Person {
		"Vincent",
		23,
		3.2,
	}

	updateName(&me, "Julius")
	updateAge(&me, 40)

	fmt.Printf("Hello my name is %v and i'm %v, i've scored %v the last time.\n", me.Name, me.Age, me.Score)
}

func updateName(person *Person, newName string) {
	person.Name = newName
}

func updateAge(person *Person, newAge int) {
	(*person).Age = newAge
}
