package main

import (
	"fmt"
)

func main() {

	// for loop example
	//forloopexample()

	//nested for loop example
	//nestedforloopexample()

	//anotherforloop()

	// infinite for loop example
	//infiniteforloop()

}

func forloopexample() {

	// Refer to https://gobyexample.com/for

	//for loop example
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

func nestedforloopexample() {

	//nested for loop example

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		for j := 0; j <= 3; j++ {
			fmt.Println(j)
		}
	}
}

func anotherforloop() {
	fmt.Println("Another for loop example")
	a := 1
	b := 5

	for a <= b {
		fmt.Println(a)
		a++
	}
}

func infiniteforloop() {
	i := 1
	for {
		fmt.Println(i)
		i++
	}
}
