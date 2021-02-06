package main

import(
	"fmt"
)

func main(){
n1:= foo()
n2,msg := bar()

fmt.Println(n1)
fmt.Println(n2)
fmt.Println(msg)
}

func foo()int{
	return 10
}

func bar()(int,string){
	return 20,"Hello World!"
}