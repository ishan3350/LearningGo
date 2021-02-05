package main

import (
	"fmt"
	"strconv"
)

func main(){
  // fmt.Println("Hello World!")
 
  // defer will delay the execution of function. It will execute the defered the code and it will call before exiting main function

  // This will execute at the end
  defer sayhello()

  displaymessage("Hello World!")

  //We can call function as many times as we want

  displaymessage("Hello again!")

  fmt.Println(add(10,20))

  resultmsg,result := subtract(20,10)

  fmt.Println(resultmsg)
  fmt.Println(result)

  // Variadic function example (We can pass as many arguments as we want)

  
  // unfurling slice or passing slice to variadic parameter

  x := []int{10,20,30,40,50,60}

  // As you know varidaic parameter is slice of specified type but you can't directly pass slice as argument. You have to add ... operator at the end
  addall(x...)


}

func sayhello(){
  fmt.Println("Hello from function")
}

// function with parameter

func displaymessage(msg string){
  fmt.Println(msg)
}

// function returning value

func add(a int, b int) int{
  return a+b
}

// returning multiple values

func subtract(a int, b int)(int, string){
  ans := "Subtraction of b from a: "+strconv.Itoa(a-b)
  return a-b,ans
}

// function with unlimited parameters (Variadic parameter)

func addall(nums...int){

  // variadic paramter is considered as slice of specified type
  result := 0
  for i:= range nums{
    result = result + nums[i]
    fmt.Println(result)
  }
}

