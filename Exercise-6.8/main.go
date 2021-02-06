package main

import (
	"fmt"
)

func main() {

	msg := displaymsg()
	msg()

}

// returning function
func displaymsg() func() {
	fmt.Println("Hello World!")
	return func() {
		fmt.Println("Hello World from function")
	}
}
