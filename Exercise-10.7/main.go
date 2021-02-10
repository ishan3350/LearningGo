package main

import (
	"fmt"
)

func main(){

	c := make(chan int)

	go func(){
		for i:=0;i<10;i++{
			for j:=0;j<10;j++{
				c <-j
			}
		}
		close(c)
	}()

	for val := range c{
		fmt.Println(val)
	}
}