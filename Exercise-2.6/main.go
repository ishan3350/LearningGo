package main

import "fmt"

const(
	year1 = 2017 + iota
	year2 = 2017 + iota
	year3 = 2017 + iota
	year4 = 2017 + iota
)

func main(){
	fmt.Println(year1)
	fmt.Println(year2)
	fmt.Println(year3)
	fmt.Println(year4)

}