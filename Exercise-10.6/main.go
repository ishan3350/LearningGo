package main

import (
	"fmt"
)

func main(){
	c := make(chan int)
	go addNumbers(c)

	for val := range c{
		fmt.Println(val)
	}


}

func addNumbers(c chan int){
	for i:=0;i<100;i++{
		c <- i
	}
	close(c)
}