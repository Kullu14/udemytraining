package main

import "fmt"

func main() {
	for i := 65; i < 123; i++ {
		fmt.Printf("%d -> %b -> %x -> %q\n", i, i, i, i)
	}
}
