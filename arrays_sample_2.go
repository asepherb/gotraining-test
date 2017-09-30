package main

import "fmt"

// ./example2.go:21: cannot use four (type [4]int) as type [5]int in assignment
func main() {

  var five [5]int

  four := [4]int{10, 20, 30, 40}

  five = four

  fmt.Println(four)
  fmt.Println(five)
}
