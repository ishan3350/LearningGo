package main

import(
	"fmt"
)

func main(){

	c := make(chan int)

	//send 

	go func(){
		for i:=0;i<=100;i++{
			c<-i
		}
		// closing channel
		close(c)
	}()

	// receive value 

	for value := range c{
		fmt.Println(value)
	}

}