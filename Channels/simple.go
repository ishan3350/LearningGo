package main

import (
	"fmt"
)

func main(){
	
	// creating simple channel

	c := make(chan int)

	go func(){
		c <- 40
	}()

	// putting and pulling is happening at the same time so it will work
	fmt.Println(<-c)
}