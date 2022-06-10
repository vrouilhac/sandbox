package utils

import (
	"fmt"
)

func GetHelloWorld() string {
	return "Hello World!"
}

func AskInput(label string) string {
	fmt.Printf("(%v)", label)	
	fmt.Printf(" > ")
	var value string
	fmt.Scan(&value)
	return value
}

