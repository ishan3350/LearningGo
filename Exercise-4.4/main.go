package main

import(
	"fmt"
)

func main(){


	x := []int{42,43,44,45,46,47,48,49,50,51}

	// appending 52 to the existing slice
	x = append(x,52)

	fmt.Println(x)

	a := []int{56,57,58,59,60}
	// appending two slices together
	x = append(x,a...)

	fmt.Println(x)
	
}