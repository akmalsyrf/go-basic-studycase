package main

import "fmt"

type Man struct {
	Name string
}

// func (m Man) Married() { //behavior secara umum, cobalah
func (m *Man) Married() {
	m.Name = "Mr. " + m.Name

	fmt.Println("Di dalam method :", m.Name)
}

func main() {
	panjul := Man{"Panjul"}
	panjul.Married()

	fmt.Println("Diluar method :", panjul.Name)
}
