package main

import (
	"fmt"
)

func main(){

	str1 := []string{"James","Bond","Shaken","Not stirred"}
	str2 := []string{"Miss","Moneypenny","Helloooooo","James"}

	str := [][]string{str1,str2}

	
	for i,data := range str{
		fmt.Println("Record: ",i)
		for ii,idata := range data{
			fmt.Printf("\nIndex position: %d\t data: %v",ii,idata)
		}
	}
}