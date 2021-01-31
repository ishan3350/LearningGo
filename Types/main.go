package main

import (
	"fmt"
)

func main() {
	fmt.Println("Types example")

	// type boolean example
	//boolexample()

	// numeric types
	numerictypes()
}

func boolexample() {

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
	fmt.Println(a == b)

}

func numerictypes() {

	// Refer here for different numeric types and their sizes https://golang.org/ref/spec#Numeric_types

	// compiler will automatically pick the best type for us. If you are not happy with compiler selection and then you can use var keyword and force it to use certain type
	a := 40
	b := 40.507590

	fmt.Printf("\n\nType of a: %T ",a)
	fmt.Printf("\nValue of a: %v ",a)

	fmt.Printf("\n\nType of b: %T: ",b)
	fmt.Printf("\nValue of b: %v: ",b)

	// You can't assign floting point value to type int variables. It's compiler error
	//a = b
	//fmt.Println(a)

	var x uint8 = 1

	fmt.Printf("\n\nType of x: %T ",x)
	fmt.Printf("\nValue of x: %v ",x)



}
