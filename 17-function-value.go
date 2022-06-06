package main

import "fmt"

func getGoodbye(name string) string {
	return "Goodbye " + name
}

//seperti closure di javascript
func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("Budi"))
}
