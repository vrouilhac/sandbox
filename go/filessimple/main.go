package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content, err := ioutil.ReadFile("test.txt")

	if err != nil {
		fmt.Println("SOMETHING WENT WRONG HERE DUDE")
	}

	fmt.Println(string(content))
}
