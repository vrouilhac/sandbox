package lib

import "fmt"

func SayHello(name string) {
	message := fmt.Sprintf("Hello %v!", name)	
	fmt.Println(message)
}
