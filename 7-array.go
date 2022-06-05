package main

import "fmt"

//array hanya bisa menyimpan jenis tipe data yang sama
func main() {
	var names [3]string // array of 3 strings
	names[0] = "John"
	names[1] = "Jane"

	fmt.Println(names)
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(len(names)) //panjang arraynya bukan jumlah isi arraynya

	var values = [3]int{
		80,
		90,
		100,
	}
	fmt.Println(values)
}
