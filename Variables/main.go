package main

import "fmt"

var z string

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

	// anotther way to declare variables

	var y = 40
	fmt.Println(y)

	// difference between short declaration and var declaration is you can't use short declaration outside of function. It will only work with in function

	test()
}

func test() {
	fmt.Println(z)
	z = "Hello World!"
	fmt.Println(z)
}
