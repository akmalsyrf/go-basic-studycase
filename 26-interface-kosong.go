package main

import "fmt"

func Ups() interface{} { //bisa return nilai apapun
	// return nil
	// return 1
	// return "ups"
	return true
}

func main() {
	data := Ups()
	// var data bool = Ups() // error implementasi interface kosong
	// var data interface{} = Ups() // ini tidak error
	fmt.Println(data)
}
