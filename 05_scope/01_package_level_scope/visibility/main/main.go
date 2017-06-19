package main

import (
       "fmt"
       "github.com/Kullu14/udemytraining/05_scope/01_package_level_scope/visibility/vis"
      )

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
