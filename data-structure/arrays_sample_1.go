package main

import "fmt"

func main() {

  var strings [5]string
  strings[0] = "apple"
  strings[1] = "banana"
  strings[2] = "orange"
  strings[3] = "grape"
  strings[4] = "plum"

  //iterate over array using range
  for i, fruit := range strings {
    fmt.Println(i, fruit)
  }

  numbers := [4]int{10, 20, 30, 40}
  count := len(numbers)
  for i:= 0; i < count; i++ {
    fmt.Println(i, numbers[i])
  }
}
