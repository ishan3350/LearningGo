package main

import(
	"fmt"
)

func main(){

	displayyears()

}

func displayyears(){
	bornyear := 1999
	currentyear := 2021

	for bornyear<=currentyear{
		fmt.Println(bornyear)
		bornyear++
	}
}