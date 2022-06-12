package main

//bisa run dengan cara :
//go run 36-package-flag.go -host=localhost -username=root -password=root -number=30

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "Put your database username")
	password := flag.String("password", "", "Put your database password")

	//cara lain
	var number *int = flag.Int("number", 0, "Put your number")

	flag.Parse()

	fmt.Println("host:", *host)
	fmt.Println("username:", *username)
	fmt.Println("password:", *password)
	fmt.Println("number:", *number)
}
