package main

import(
	"fmt"
)

func main(){
	defer fmt.Println("This will print at the end of program")
	displaymessage()
}

func displaymessage(){
	fmt.Println("Hello World!")
}