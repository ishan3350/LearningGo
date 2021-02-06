package main

import(
	"fmt"
)

type programmer interface{
	code()
}

type juniorprogrammer struct{
	name string
	age int
}

type advancedprogrammer struct{
	name string
	age int
	experience int
}

func(p juniorprogrammer) code(){
	fmt.Println("Name: ",p.name)
	fmt.Println("Age: ",p.age)

}
func (p advancedprogrammer)code(){
	fmt.Println("Name:", p.name)
	fmt.Println("Age: ",p.age)
	fmt.Println("Experience: ",p.experience)


}
func main(){

	p1 := juniorprogrammer{
		name: "Programmer 1",
		age: 10,
	}

	p2 := advancedprogrammer{
		name: "Programmer 2",
		age: 20,
		experience: 10,
	}

	p1.code()
	p2.code()
}

