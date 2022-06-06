package main

import "fmt"

func sayHello(firstName string, lastName string) (string, int) {
	return "Hello " + firstName + " " + lastName, 123 // return multiple value
}

func main() {
	result, _ /*return kedua diignore*/ := sayHello("John", "Doe")
	fmt.Println(result)
	// fmt.Println(integer)
}
