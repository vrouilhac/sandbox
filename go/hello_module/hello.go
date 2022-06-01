package hello_module

import "fmt"

func SayHello(name string) string {
	message := fmt.Sprintf("Hello %v !", name);	
	return message
}
