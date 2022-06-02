package main

import "fmt"

func main() {
	deferLIFO()
	deferFunction()
}

func getAge() int {
	return 23
}

func smallSwitch() {
	switch age := getAge(); age {
		case 20:
			fmt.Println("You boy!")
		case 21:
		case 22:
			fmt.Println("Not much")
		case 23:
			fmt.Println("My age dude!")
		default:
			fmt.Println("Cannot guess")
	}
}

func specialSwitch() {
	switch {
		case getAge() < 12:
			fmt.Println("You kido")
		case getAge() < 18:
			fmt.Println("You teenager")
		case getAge() < 22:
			fmt.Println("You grown up")
		case getAge() < 45:
			fmt.Println("You man")
		default:
			fmt.Println("You wise man")
	}
}

func deferFunction() bool {
	defer fmt.Println("I will print once this function return")

	fmt.Println("I print inside this function")

	return true
}

func deferLIFO() {
	fmt.Println("Counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("Finish Counting")
}
