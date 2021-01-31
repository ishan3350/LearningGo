package main

import "fmt"

func main(){

	display()

}

func display(){
	for i:=65;i<=90;i++{
		for j:=1;j<=3;j++{

			fmt.Printf("\n%U\t %q ",i,i)
		}
	}
}

