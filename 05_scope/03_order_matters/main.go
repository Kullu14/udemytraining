package main

import "fmt"

func main() {
	//fmt.Println(x)   //Order matters this won't compile
	x := 34
	fmt.Println(x);
	fmt.Println(y);
}

var y = 45  //package scope order doesn't matter
