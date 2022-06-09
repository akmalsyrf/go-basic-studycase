package main

import "fmt"

func main() {
	counter := 0
	inrement := func() {
		fmt.Println(counter)
		counter++ //closure
	}
	inrement()
	inrement()
	fmt.Println(counter)
}
