package main

import "fmt"

var x = 10

func decrement () int {
	x--
	return x
}

func main() {
	fmt.Println(decrement())
	fmt.Println(decrement())
}
