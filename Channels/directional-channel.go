package main

import (
	"fmt"
)

func main(){

	// send only channel
	c := make(chan <- int,2 )

	c <- 40
	c <- 45

	// it's send only channel and we are trying to receive value from channel

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	fmt.Println("-------------------------------------")
	fmt.Printf("%T\n",c)


	// receive only channel

	//ch := make(<- chan int,2)

	// We are trying to send value in receive only channel
	// ch <- 40
	// ch <- 45


}