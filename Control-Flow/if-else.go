package main

import (
	"fmt"
)

func main() {

	//simpleifelse()

	//anotherifelse()

	elseifexample()

}

func simpleifelse() {
	a := 10
	b := 20

	if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("b is greater than a ")
	}
}

// we can intialize variable in if block and it will limit within that block
func anotherifelse() {
	if a := 10; a < 5 {
		fmt.Println("a is less than 5")
	} else {
		fmt.Println("a is greater than 5")
	}
}

func elseifexample() {
	a := 10
	b := 20

	if a < b {
		fmt.Println("a is greater than b")
	} else if a > b {
		fmt.Println("a is greater than b")
	} else {
		fmt.Println("a is equal to b")
	}
}
