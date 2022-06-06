package main

import "fmt"

func getCompleteName() (firstName, lastName string) {
	firstName = "John"
	lastName = "Doe"
	// return firstName, lastName //ini bisa juga
	return
}

func main() {
	a, b := getCompleteName()
	fmt.Println(a, b)
}
