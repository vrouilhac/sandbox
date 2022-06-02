package hello

import (
	"fmt"
	"time"
	"math"
	"math/rand"
)

const isDev = false

func Resolve() {
	fmt.Println(time.Now())
	fmt.Println(rand.Intn(20))
	fmt.Println(math.Sqrt(9))
	fmt.Println(add(3, 5))
	x := "World!"
	y := "Hello"
	swapedX, swapedY := swap(x, y)
	message := fmt.Sprintf("%v %v", swapedX, swapedY)
	fmt.Println(message)

	a, z := split(20)

	if isDev {
		fmt.Printf("a: %v, z: %v\n", a, z)
	}
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(input int) (x, y int) { // (x, y int) are named import
	x = input / 2
	y = input - x
	return
}

// Go's basic type https://go.dev/tour/basics/11
