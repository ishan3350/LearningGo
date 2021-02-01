package main

import(
	"fmt"
)

func main(){

	/*
	Need below output using slicing slice

	[42 43 44 45 46]
	[47 48 49 50 51]
	[44 45 46 47 48]
	[43 44 45 46 47]

*/
	x := []int{42,43,44,45,46,47,48,49,50,51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}