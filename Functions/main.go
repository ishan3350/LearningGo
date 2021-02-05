package main

import(
  "fmt"
)

func main(){
  // fmt.Println("Hello World!")

  sayhello()

  displaymessage("Hello World!")

  //We can call function as many times as we want

  displaymessage("Hello again!")
}

func sayhello(){
  fmt.Println("Hello from function")
}

// function with parameter

func displaymessage(msg string){
  fmt.Println(msg)
}
