package main

import "fmt"

func random() interface{} {
	return "OK"
}

func main() {
	result := random()
	resultString := result.(string) // type assertion, jika result sudah dipastikan string
	fmt.Println(resultString)

	// resultInt := result.(int) // panic
	// fmt.Println(resultInt)

	//lebih recomended menggunakan switch case
	switch result.(type) {
	case string:
		fmt.Println("Value", result, "is string")
	case int:
		fmt.Println("Value", result, "is int")
	default:
		fmt.Println("unknown type")
	}
}
