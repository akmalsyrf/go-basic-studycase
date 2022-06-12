package database

import "fmt"

var connection string

func init() {
	fmt.Println("Initializing database...")
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
