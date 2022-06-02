package main

import (
	"fmt"
)

const (
	PI = 3.1415
	NAME = "Vincent"
	LEARNING_SECTION = "infinity" // Could be "for" | "while" | "infinity"
)

func main() {
	fmt.Println(fmt.Sprintf("Hello %v, did you know pi is %v ?", NAME, PI))

	if LEARNING_SECTION == "for" {
		ForLoop()
	}

	if LEARNING_SECTION == "while" {
		sum := WhileLoop(5)

		fmt.Printf("Sum(%v)\n", sum)
	}

	if LEARNING_SECTION == "infinity" {
		// Infinity()
		if v := UselessFunctionDude(); v > 0 {
			fmt.Println("This is the right case !")
		}
	}
}

func ForLoop() {
	count := 0

	for i := 0; i < 10; i++ {
		count += i
		fmt.Println(fmt.Sprintf("Value of i is %v", i))
	}

	fmt.Printf("Count is {%v}\n", count)

	c := 0
	for ; count < 100 ; {
		c++ // the badass who code c++ inside go o_O
		count += count
	}

	fmt.Printf("Count is {%v}\n", count)
	fmt.Printf("C++ is {%v} (Very good and by good i mean hard!)\n", c)
}

func UselessFunctionDude() int {
	return 5
}

func WhileLoop(step int) int {
	sum := 0
	for sum < 100 {
		sum += step
	}
	return sum
}

func Infinity() {
	count := 0

	for {
		if count == 5 {
			break
		}

		fmt.Println("Not there yet")
	}

	fmt.Printf("count(%v)", count)
}
