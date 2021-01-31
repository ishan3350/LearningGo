package main

import(
	"fmt"
)

func main(){
	switchexample()
}

func switchexample(){

	switch{
	case 5==5:
		fmt.Println("5==5 is true")
		fallthrough
	case 10==10:
		fmt.Println("10==10 is true")
	default:
		fmt.Println("No conditions are matching")
	}
}