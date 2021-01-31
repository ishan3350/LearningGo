package main

import (
	"fmt"
)

func main() {
	simpleswitch()
}

func simpleswitch() {
	num := 5

	switch num {
	case 1:
		fmt.Println("number is one")
		fmt.Println("number is 1")

	case 2:
		fmt.Println("number is two")

	case 5:
		fmt.Println("number is five")

	default:
		fmt.Println("Number is not one, two or five")
	}

}
