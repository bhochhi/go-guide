package main

import (
	"fmt"

	"../util"
)

func add(x, y int) int {
	return x + y
}

func main() {
	java := "hello"
	fmt.Println("Hi there!! ", java)
	fmt.Println(add(2, 4))
	fmt.Println(util.Add(5, 11))
}
