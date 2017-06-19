package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)
	foo()
}

func foo() {
	fmt.Println(x)  //this line won't work
}
