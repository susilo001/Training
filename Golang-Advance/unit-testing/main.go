package main

import "fmt"

func GetHelloMessage() string {
	return "Hello Hello"
}
func main() {
	message := GetHelloMessage()

	fmt.Println(message)
}
