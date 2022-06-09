package main

import "fmt"

func factorialLoop(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorialRecursive(n-1)
}

func main() {
	loop := factorialLoop
	fmt.Println(loop(5))
	fmt.Println(factorialRecursive(5))
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
