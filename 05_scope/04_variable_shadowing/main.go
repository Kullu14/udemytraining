package main

import "fmt"

func add (x int) int {
	return (x + x)
}

func main() {
	add := add(10)
	fmt.Println(add)
}


// bad coding practise
