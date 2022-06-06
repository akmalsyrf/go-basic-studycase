package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	//for dengan statement
	slice := []string{"Eko", "Joko", "Budi"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	//for range, biasanya untuk iterasi array, slice, map
	names := []string{"Eko", "Joko", "Budi"}
	for index, value := range names {
		fmt.Println(index, value)
	}
}
