package main

import(
	"fmt"
)
// this example is not working as send and receive is not performing at the same time

func main(){
	fmt.Println("Hello World!")

	// creating channel

	c := make(chan int)

	// putting value in channel
	//send and receive needs to be done at the same time or channel will block

	c <- 40


	fmt.Println(<-c)
}