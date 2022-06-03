package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	arr := make([]string, 5, 10)
	// fmt.Println(arr)

	arr[0] = "Bonjour"
	// arr[7] = "Vincent" // what is the case where we can to call make
	// with length < capacity ?

	// fmt.Println(arr)

	array := []int{1, 2, 3}

	array2 := append(array, 4, 5, 6)

	fmt.Println(array, array2)

	PrintArray(array2)
}

func PrintArray(arr []int) {
	for _, v := range arr {
		fmt.Println(v)
	}
}
