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

type expertprogrammer struct{
	name string
	age int
	experience int
	numawards int
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

func (p expertprogrammer)code(){
	fmt.Println("Name:", p.name)
	fmt.Println("Age: ",p.age)
	fmt.Println("Experience: ",p.experience)
	fmt.Println("Number of awards: ",p.numawards)

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

	p3 := expertprogrammer{
		name: "Programmer 3",
		age: 25,
		experience: 21,
		numawards: 5,
	}

	p1.code()
	p2.code()
	p3.code()
}

