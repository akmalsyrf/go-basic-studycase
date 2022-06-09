package main

import "fmt"

//bisa dianggap seperti class di bahasa pemrograman oop, kita mengelompokan prototype data yg sejenis dalam satu struct

type Customer struct {
	Name, Address string
	Age           int
}

//method struct
func (customer Customer) sayHello() {
	fmt.Println("Hello, my name is", customer.Name)
}

func main() {
	var panjul Customer
	panjul.Name = "Panjul"
	panjul.Address = "Jl in aja dulu siapa tau cocok"
	panjul.Age = 20

	fmt.Println(panjul)
	panjul.sayHello()

	//atau bisa juga dengan menggunakan literal
	joko := Customer{
		Name:    "Joko",
		Address: "Jl in aja dulu siapa tau cocok",
		Age:     20,
	}
	fmt.Println(joko)
	joko.sayHello()

	budi := Customer{"Budi", "Jl in aja dulu siapa tau cocok", 20}
	fmt.Println(budi)
	budi.sayHello()
}
