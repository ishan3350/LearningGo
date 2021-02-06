package main

import(
	"fmt"
)
func main(){
	displayData(10)
	displayData("Hello World")
}

// it will allow us to pass anytype of value
func displayData(a interface{}){
	fmt.Println(a)
}