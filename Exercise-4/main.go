package main

import (
	"fmt"
)
type hiint int

var x hiint

func main(){

fmt.Println("Value of x: ", x)
fmt.Printf("Type of x %T:",x )
x = 42
fmt.Println(x)

}