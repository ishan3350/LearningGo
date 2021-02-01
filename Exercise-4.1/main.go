package main

import(
	"fmt"
)

func main(){

	//creating array using string literal
	a := [5]int{10,20,30,40,50}

	// looping through array using for and range
	for val := range a{
		fmt.Println(val)
	}

	fmt.Printf("Type of array %T",a)
}