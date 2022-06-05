package main

import (
	"fmt"
)

func main() {
	var name string
	// var age int = 24 // error jika variabel tidak digunakan
	hobby := "programming" //cara lain untuk deklarasi variabel

	name = "John Doe"
	name = "Jane Doe" //bisa di asign value baru
	// name = 24      //error, karena string tidak bisa di assign nilai lain

	fmt.Println(name)
	fmt.Println(hobby)
}
