package main

import "fmt"

func main() {
	fmt.Printf("%d - %b - %x\n", 21, 21, 42)
	fmt.Printf("%d - %b - %X\n", 21, 21, 42)
	fmt.Printf("%d - %b - %#x\n", 21, 21, 42)
	fmt.Printf("%d - %b - %#X\n", 21, 21, 42)
	fmt.Printf("%d \t %b \t %x\n", 21, 21, 42)
}
