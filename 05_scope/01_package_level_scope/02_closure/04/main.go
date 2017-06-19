package main

import "fmt"

func wrapper() func() int {
	x := 0 
	return func() int {
		x--
		return x
	}
}

func main() {
	decrement := wrapper()
	fmt.Println(decrement())
	fmt.Println(decrement())
}
