package main

import "fmt"

func main() {
	name := "Budi"

	if name == "Budi" {
		fmt.Println("Halo Budi")
	} else if name == "Eko" {
		fmt.Println("Halo Eko")
	} else {
		fmt.Println("Halo yang lain")
	}

	//if dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
