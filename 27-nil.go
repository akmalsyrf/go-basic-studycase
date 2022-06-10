package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("Panjul")
	if person == nil {
		fmt.Println("Person kosong")
	} else {
		fmt.Println(person)
	}
}
