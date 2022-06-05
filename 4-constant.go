package main

import "fmt"

func main() {
	const name string = "John Doe" // deklarasi nilai harus diisi diawal
	// name = "Jane Doe" //error, karena const tidak bisa di assign value baru
	// name = 24         //error, karena const tidak bisa di assign nilai lain

	//deklarasi multiple const
	const (
		firstName = "Budi" // tidak ada error meskipun variabel tidak digunakan
		lastName  = "Wahyudiputra"
	)
	fmt.Println(name)
}
