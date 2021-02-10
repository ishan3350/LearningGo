package main

import(
	"fmt"
)
type person struct{
	first string
	last string
	age int
}

type human interface{
	speak()
}
func (p *person) speak(){
	fmt.Println(p.first, "is able to speak")
}

func saySomething(h human){
	h.speak()
}
func main(){

	p1 := person{
		first: "test",
		last: "test2",
		age: 30,
	}

	saySomething(&p1)

}