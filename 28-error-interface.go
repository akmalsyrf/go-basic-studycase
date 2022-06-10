package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) { // interface error
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0") //package errors
	}
	return nilai / pembagi, nil
}
func main() {
	var nilai int
	var pembagi int
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)
	fmt.Print("Masukkan pembagi: ")
	fmt.Scan(&pembagi)
	hasil, err := Pembagian(nilai, pembagi)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hasil: ", hasil)
	}
}
