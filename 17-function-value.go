package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("Budi"))
}
