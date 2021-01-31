package main

import (
	"fmt"
)
type hiint int

var x hiint

var y int

func main(){

fmt.Println("Value of x: ", x)
fmt.Printf("Type of x %T:",x )
x = 42
fmt.Println(x)
// execrcise 05

y = int(x)
fmt.Println("Value of y: ",y)
fmt.Printf("Type of y: %T: ",y)

}