package main

import(
	"fmt"
)

func main(){

	displayremainder()

}

func displayremainder(){
	for i:=10;i<=100;i++{
		fmt.Printf("\n%v\t%d",i,i%4)
	}
}
