package main

import "fmt"

func main() {
	type NoKtp string
	type Married bool

	var myKtp NoKtp = "123456789"
	var myMarried Married = false

	fmt.Println(myKtp)
	fmt.Println(NoKtp("777777777"))
	fmt.Println(myMarried)
}
