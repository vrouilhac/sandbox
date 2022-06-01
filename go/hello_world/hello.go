package main

import (
	"fmt"

	"vrouilhac/hello_module"
)

import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
	fmt.Println(hello_module.SayHello("Vincent"))
}
