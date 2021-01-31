/*
1. Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS "x" and "y" and "z"

a. 42
b. "James Bond"
c. true

2. Now print the values stored in those variables using

a. a single print statement
b. multiple print statements

*/

package main

import (
	"fmt"
)

func main() {

	x := 42
	y := "James Bond"
	z := true

	// printing values of variables using single print statement
	fmt.Print("Value of x: ", x, " Value of y: ", y, " Value of z: ", z,"\n")

	// printing values of variables using multiple print statements

	fmt.Print("Value of x: ",x,"\n")
	fmt.Print("Value of y: ",y,"\n")
	fmt.Print("value of z: ",z,"\n")

}
