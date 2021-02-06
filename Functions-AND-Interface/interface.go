package main

import(
	"fmt"
)

type person struct{

	name string
	age int
}

type child struct{
	name string
	age int
	canwalk bool
}

type human interface{
	// any other type that has method is also human type
	// value can be of more than one type
	speak()
	
}
func main(){

	// interface example

	p1 := person{
		name:"test",
		age:30,
	}
	
	fmt.Println(p1)
	p1.speak()

	c1 := child{
		name:"test2",
		age: 30,
		canwalk: true,
	}

	fmt.Println(c1)

	c1.speak()

	// as you can see display message accept both types child and person as speak method is attached to all three human, child and person
	displaymessage(p1)

	displaymessage(c1)
}

func (p person) speak(){
	fmt.Println(p.name, "is able to speak")
}

func (c child) speak(){
	fmt.Println(c.name, "child is able to speak")
}


// interface allows value of more than one type
func displaymessage(h human){
	fmt.Println("Displaying simple message from human")
}

