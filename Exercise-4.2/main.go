package main

import(
	"fmt"
)

func main(){

	x := []int{1,2,3,4,5}
	
	for val := range x{
		fmt.Println(val)
	}

	// displaying data type of slice
	fmt.Printf("Type of slice: %T",x)
	


}