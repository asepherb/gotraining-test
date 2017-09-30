package main

import "fmt"

func main() {

  var n int

  //declare anonymous functions and call it
  func(){
    fmt.Println("direct :", n)
  }()

  //declare anonymous function and assign to variable
  f := func(){
    fmt.Println("variable : ", n)
  }

  //call it
  f()

  //defer the call to the anonymous function till after main returns
  defer func() {
    fmt.Println("defer 1: ", n)
  }()

  n = 3
  f()

  defer func() {
    fmt.Println("defer 2: ", n)
  }()
}
