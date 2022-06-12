package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	hostname := os.Hostname
	fmt.Println(hostname)

	gopath := os.Getenv("GOPATH")
	fmt.Println(gopath)
}
