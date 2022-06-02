package main

import "fmt"

func main() {
	var username string;
	
	fmt.Println("Hi, what's your name ? ")

	fmt.Scan(&username)

	fmt.Printf("Hello %v, nice to meet you\n", username)
}
