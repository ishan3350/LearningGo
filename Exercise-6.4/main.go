package main

import (
	"fmt"
	"strconv"
)

type person struct{
	first string
	last string
	age int
}
func (p person)speak(){
fmt.Println("Name: "+ p.first+" "+p.last+" is "+ strconv.Itoa(p.age)+" years old")
}
func main(){

	p1 := person{
		first: "Test",
		last: "Test2",
		age: 25,
	}

	p1.speak()
}