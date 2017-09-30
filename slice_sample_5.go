package main

import "fmt"

func main() {

  x := make([]int, 7)

  for i := 0; i < 7; i++ {
    x[i] = i * 100
  }

  twoHundred := &x[1]

  x = append(x, 800)

  x[1]++

  fmt.Println("pointer:", *twoHundred, "element", x[1])
}
