package main

import "fmt"

//map adalah tipe data yang bisa menampung banyak nilai dengan key-value
func main() {
	person := map[string]string{
		"name": "Eko",
		"age":  "20",
	}

	person["hobby"] = "Programming"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["hobby"])

	//function di map
	fmt.Println(len(person))
	// fmt.Println(cap(person)) // tidal ada cap di map
	fmt.Println(make(map[string]string)) //membuat map baru

	//delete
	delete(person, "age")
	fmt.Println(person)
}
