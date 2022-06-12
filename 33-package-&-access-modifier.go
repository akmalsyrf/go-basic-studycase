package main

import (
	"aKoding/go-basic-studycase/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Budi")
	fmt.Println(result)

	result2 = helper.sayGoodbye("Budi")
	fmt.Println(result2) //tidak keluar karena tidak bisa diakses di package lain
}
