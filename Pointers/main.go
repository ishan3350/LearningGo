package main

import (
	"fmt"
)

func main() {
	a := 40
	fmt.Println(a)

	// getting memory address of variable
	fmt.Println(&a)

	// getting type (*int)
	fmt.Printf("%T", &a)

	// & person will give you adderess of variable in memory where it is stored

	var b *int = &a
	fmt.Println(b)

	//accessing pointer value
	fmt.Println(*b)

	// *& will also give value at that address
	fmt.Println(*&a)

	// changing value using pointer
	*b = 20
	fmt.Println(a)

	x := 10
	fmt.Println(x)
	changevalue(&x)
	fmt.Println(x)

}

// pointer as argument in function
func changevalue(a *int) {
	*a = 20

}
