package main

import(
	"fmt"
)

func main(){

	//simple anonymous function without any parameter
	func(){
		fmt.Println("Hello from anonymous function")
	}()

	//anonymous function with paramter
	func(a int){
		fmt.Println(a*2)
	}(10)

	// stroing functiion in indetifier
	multi := func (x int, y int){
		fmt.Println(x*y)
	}
	multi(10,5)

	// anonymous function returning value

	division := func(a float64, b float64)(float64){
		return (a/b)
	}

	fmt.Println(division(50,2))

}
