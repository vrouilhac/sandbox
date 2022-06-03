package main

import (
	"fmt"
	// "time"
)

func main() {
	age := 23

	var p *int
	p = uselessFunctionToGetThePointerOfAVariableNotTooLong(age)

	fmt.Println(p)

	add(p, 20)

	fmt.Println(age)
}

func add(value *int, ajout int) {
	*value += ajout
	fmt.Println(*value)
}

// That was a good lesson to learn, value is declared here so the pointer to value
// is not the pointer to age, age has been duplicated to 'value' therefor i return
// a pointer to a variable that is being garbage collected ? 
func uselessFunctionToGetThePointerOfAVariableNotTooLong(value int) *int {
	fmt.Println("Yes i also print something dude")
	fmt.Println(&value)
	return &value
}
