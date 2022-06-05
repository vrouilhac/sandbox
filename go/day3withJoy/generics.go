package main

import (
	"fmt"	
)


func Compare[T comparable](x T, y T) bool {
	if x == y {
		return true
	}

	return false
}

func main() {
	fmt.Printf("%v == %v\t: %v\n", 3, 2, Compare(3, 2))
	fmt.Printf("%v == %v\t: %v\n", "Vincent", "Vincent", Compare("Vincent", "Vincent"))
	fmt.Printf("%v == %v\t: %v\n", 3, 3, Compare(3, 3))
	fmt.Printf("%v == %v\t: %v\n", 2, 5, Compare(2, 5))
	fmt.Printf("%v == %v\t: %v\n", "Vincent", "Coucou", Compare("Vincent", "Coucou"))
}
