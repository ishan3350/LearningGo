package main

import(
	"fmt"
)

func main(){

	// for loop example
forloopexample()

//nested for loop example
//nestedforloopexample()
	
}

func forloopexample(){
	for i:=0;i<=10;i++{
		fmt.Println(i)
	}
}
func nestedforloopexample(){

	for i:=0;i<=10;i++{
		fmt.Println(i)
		for j:=0;j<=3;j++{
			fmt.Println(j)
		}
	}
}