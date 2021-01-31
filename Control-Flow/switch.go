package main

import (
	"fmt"
)

func main() {
	//simpleswitch()
	anotherswitchstatemnet()
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

// fallthrough keyword tells compiler to check next case even if they found that case matching condition or expectations
func anotherswitchstatemnet() {

	switch {
	case 1 == 1:
		fmt.Println("this should print")
		fallthrough

	case 2 == 2:
		fmt.Println("this should print")
		fallthrough

	case 3 == 3:
		fmt.Println("this should print")
		fallthrough

	case 4 == 4:
		fmt.Println("this should print")

	case 5 == 5:
		fmt.Println("This should not print")
	}

}
