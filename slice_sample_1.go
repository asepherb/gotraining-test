package main

import "fmt"

func main() {

  slice := make([]string, 5)
  slice[0] = "apple"
  slice[1] = "banana"
  slice[2] = "pulm"
  slice[3] = "orange"
  slice[4] = "grape"

  fmt.Println(slice)
}
