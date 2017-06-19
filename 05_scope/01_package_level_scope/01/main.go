package main

import "fmt"

var x int = 23

func main() {
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x + x)
}
