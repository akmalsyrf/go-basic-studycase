package main

import "fmt"

//destructuring arguments kalo di javascript
func sumAll(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	fmt.Println(sumAll(1, 2, 3, 4, 5))
	fmt.Println(sumAll(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(sumAll(slice...))
}
