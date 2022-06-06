package main

import "fmt"

func main() {
	name := "Budi"

	switch name {
	case "Budi":
		fmt.Println("Budi")
	case "Eko":
		fmt.Println("Eko")
	default:
		fmt.Println("Tidak diketahui")
	}

	//switch dengan short statement
	switch length := len(name); length > 5 { //scope nya hanya di switchnya saja
	case true:
		fmt.Println("Nama terlalu panjang")
	case false:
		fmt.Println("Nama sudah benar")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 5:
		fmt.Println("Nama terlalu panjang")
	case length < 5:
		fmt.Println("Nama sudah benar")
	}
}
