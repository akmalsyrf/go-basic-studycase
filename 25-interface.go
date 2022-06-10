package main

import "fmt"

type HasName interface {
	GetName() string
}

type Person struct { // bayangkan ini class Person
	Name string
}

type Animal struct { //bayangkan ini class Animal
	Name string
}

func (p Person) GetName() string { //method dari class Person
	return p.Name
}
func (a Animal) GetName() string { //method dari class Animal
	return a.Name
}

func SayHello(h HasName) {
	fmt.Println("Hello, my name is", h.GetName())
}

func main() {
	var panjul Person
	panjul.Name = "Panjul"

	SayHello(panjul)

	cat := Animal{Name: "Cat"}
	SayHello(cat)
}
