package main

import(
	"fmt"
)

func main(){
	//buffer channel. It will allow one value to sit regardless of something is ready to pull it off or not
	c := make(chan int,1)

	c <- 40

	fmt.Println(<-c)

	
}