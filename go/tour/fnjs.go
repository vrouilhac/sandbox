package main

import "fmt"

func main() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(computeAdd(add, 4, 5))
}

func computeAdd(fn func(x, y int) int, x int, y int) int {
	return fn(x, y)
}
