package main

import "fmt"

func main() {

	// You can't use keyword for variable names. Refer to https://golang.org/ref/spec#Keywords

	//declaring variable and assigning a value
	x := 10
	fmt.Println(x)

	// changing value of existing value

	x = 20
	fmt.Println(x)

	// peroforming operation and assigning value to the variable (statement)
	// 10 and 20 are operand and + is operator and 10+20 is instruction
	x = 10 + 20
	fmt.Println(x)

}
