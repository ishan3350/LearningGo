package main

import (
	"fmt"
)

func main() {

	c := make(chan int,1)
	// solution 1
	// go func() {
	// 	fmt.Println(<-c)
	// }()

	// c <- 42

	// solution 2
	c <- 42
	fmt.Println(<-c)

	}