package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp(value int) {
	// fmt.Println("Run app")
	// result := 10 / value // error karena 10 / 0 tidak terhingga, sehingga logging tidak dieksekusi
	// fmt.Println("Result", result)
	// logging()

	defer logging() // logging tetap dieksekusi, karena defer function akan dieksekusi diakhir method meskipun ada panic
	result := 10 / value
	fmt.Println("Result", result)
	logging()
}

func main() {
	runApp(0)
}
