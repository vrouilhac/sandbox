package main

import (
	"fmt"
)

func main() {
	input := 139.0
	round := 100
	output := Sqrt(input, round)

	fmt.Printf("input: {%g}\noutput: {%g}\n", input, output)
	fmt.Printf("input: {%g}\noutput: {%g}\n", input, output*output)
}

func Sqrt(x float64, round int) (output float64) {
	output = 1.0

	for i := 1; i <= round; i++ {
		output -= (output*output - x) / (2.0 * output)
	}

	return
}
