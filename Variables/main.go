package main

import "fmt"

var z string

// creating type
type hiint int

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

	typeexplained()

	zerovalue()

	exampleofsprintf()

	typeexample()
}

func test() {
	fmt.Println("Printing the value of z" + z)
	fmt.Printf("Printing the type of z %T", z)
	fmt.Println("")
	z = "Hello World!"
	fmt.Println(z)
}

func typeexplained() {
	var b = 30
	fmt.Printf("\nType of B: %T", b)
	fmt.Printf("\nValue of B: %v", b)

	// You can't change the type of variable once declared
	// compiler error
	// b = "Hello World"

	fmt.Println("")

	// How to use quote in variable value
	str := `"Hello World!"`
	fmt.Println(str)
}

func zerovalue() {
	var c int
	fmt.Println(c)
}

func exampleofsprintf() {

	// You can get the value of print using string print

	//let's get the binary of number

	fmt.Println("\nGet Binary value of a number\n")
	var str1 = fmt.Sprintf("%b", 50)
	fmt.Println(str1)

}

func typeexample() {

	var d hiint = 50
	fmt.Printf("Type of d %T",d)
	fmt.Println("\n")
	fmt.Printf("Value of d %v",d)

	// type conversion example

	var e int = 60

	// we can't assign the value of d to e because d is hiint type and e is int type

	//type conversion

	d = hiint(e)
	fmt.Println("")
	fmt.Printf("Value of d %v",d)

}


