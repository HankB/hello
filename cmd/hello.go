package main

import (
	"fmt"
	"hello/pkg/lib"
)

func main() {
	fmt.Println("Hello there world!")
	fmt.Println("sum:", lib.Add(2,3))
}