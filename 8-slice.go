package main

import "fmt"

//slice mirip array, yang membedakan adalah ukuran arraynya bisa bertambah dan berkurang
func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]

	fmt.Println("Nilai awal slice1 :", slice1)
	fmt.Println(len(slice1))                          //length
	fmt.Println(cap(slice1))                          //capacity
	fmt.Println("append :", append(slice1, "Kliwon")) // kapasitas slice1 penuh, maka append akan membuat array baru

	slice1[0] = "Diubah"
	fmt.Println("nilai akhir slice1 :", slice1)
	fmt.Println("array bulan :", months)

	//make slice
	newSlice := make([]string, 3, 5) // 3 panjang, 5 kapasitas
	newSlice[0] = "Alpha"
	newSlice[1] = "Beta"

	fmt.Println("make slice :", newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	//copy slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println("copy slice :", copySlice)

	//array vs slice
	//di console hasilnya sama
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}
	fmt.Println("iniArray :", iniArray)
	fmt.Println("iniSlice :", iniSlice)
}
