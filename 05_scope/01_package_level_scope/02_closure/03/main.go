package main

import "fmt"

func main() {
	x := 0
	increment := func () int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}


/* anonymous function 
   a function without name

   function expression
   assigning a function to a variable
*/
