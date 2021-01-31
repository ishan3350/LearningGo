package main

import "fmt"

func main(){
	
	var a int = 10

	fmt.Printf("Value of a in decimal %d, value of a in binary %b, and value of a in hexadecimal %X ",a,a,a)

	x := a<<1
	fmt.Printf("Value of x in decimal %d, value of x in binary %b, and value of x in hexadecimal %X ",x,x,x)


}