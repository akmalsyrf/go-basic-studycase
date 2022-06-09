package main

import "fmt"

func endApp() {
	message := recover() //menangkap error yang terjadi, dan supaya aplikasi tidak berhenti
	if message != nil {
		fmt.Println("Error dengan message:", message)
	}

	fmt.Println("selesai memanggil function")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Aplikasi error") //membuat error dengan panic
	}
	fmt.Println("Run app")
}

func main() {
	runApp(true)
	fmt.Println("Baris ini berjalan karena ada fungsi recover")
}
