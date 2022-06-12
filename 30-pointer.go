package main

import "fmt"

//pelajari lebih lanjut soal pass by value dan pass by reference, juga operator *
//disini sengaja saya comment bagian yang menggunakan pointer, supaya bisa dipahami perbedaannya

type Address struct {
	City, Province, Country string
}

func main() {
	address1 := Address{City: "Subang", Province: "Jawa Barat", Country: "Indonesia"}

	//tanpa pointer
	address2 := address1 //pass by value, address2 mengcopy address1

	//dengan pointer
	// address2 := &address1             //pass by reference menggunakan pointer, address2 mereference address1
	// var address3 *Address = &address1 //bisa juga dengan cara

	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	// fmt.Println(address3)

	//operator new
	var address4 *Address = new(Address) //menggunakan new, menghasilkan pointer kosong (kecuali jika di assign)
	address4.City = "Jakarta"
	fmt.Println(address4)
}
