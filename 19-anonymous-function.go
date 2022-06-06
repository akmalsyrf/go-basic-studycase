package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("User", name, "is blacklisted")
	} else {
		fmt.Println("User", name, "is not blacklisted")
	}
}

func main() {
	//anonimous function dengan variabel
	blacklist := func(name string) bool {
		return name == "mark"
	}

	registerUser("mark", blacklist)
	registerUser("john", blacklist)

	//cara lain deklarasi anonymous function
	registerUser("mark", func(name string) bool {
		return name == "mark"
	})
}
