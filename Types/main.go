package main

import (
	"fmt"
)

func main(){
	fmt.Println("Types example")

	// type boolean example
	boolexample()
}

func boolexample(){

	// boolean type can hold true or false value.

	var flag bool
	// checking default value for bool which is false
	fmt.Println(flag)
	
	// changing value of boolean variable

	flag = true
	
	fmt.Println(flag)

	a := 10
	b := 20
	 // value of a is not equal to b so it will return false
	fmt.Println(a==b)

}