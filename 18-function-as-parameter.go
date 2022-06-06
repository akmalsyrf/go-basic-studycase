package main

import "fmt"

type Filter func(string) string

//seperti callback di javascript, HOF
func sayHelloWithFilter(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Anjing", spamFilter)
	sayHelloWithFilter("Bambang", spamFilter)
}
