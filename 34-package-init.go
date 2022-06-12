package main

import (
	"aKoding/go-basic-studycase/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
