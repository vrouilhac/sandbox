package main

import (
	"fmt"
)

type MyError int

func (e MyError) Error() string {
	return fmt.Sprintf("Error on line %d", e)
}

func Compute(x int) (int, error) {
	if x < 0 {
		return x, MyError(x)
	}

	return x, nil
}

func main() {
	fmt.Println(Compute(2))
	fmt.Println(Compute(-2))
}
