package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "Inside the inner scope"
		fmt.Println(y)
	}	
	//fmt.Println(y) //Outside the scope
}
