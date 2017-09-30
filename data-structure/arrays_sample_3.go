package main

import "fmt"

func main() {

  fivePeople := [5]string{"andi", "apang", "abe", "eyang", "pangeran"}
  fmt.Printf("before [%s] : ", fivePeople[1])

  for i := range fivePeople {

    fivePeople[1] = "andi_become"

    if i == 1 {
      fmt.Printf("after [%s]\n", fivePeople[1])
    }
  }

  fivePeople = [5]string{"andi", "apang", "abe", "eyang", "pangeran"}
  fmt.Printf("before [%s] : ", fivePeople[1])

  for i, v := range fivePeople {

    fivePeople[1] = "andi_become"

    if i == 1 {
      fmt.Printf("v[%s]\n", v)
    }
  }

  fivePeople = [5]string{"andi", "apang", "abe", "eyang", "pangeran"}
  fmt.Printf("before [%s] : ", fivePeople[1])
  for i, v := range &fivePeople {

    fivePeople[1] = "andi_become"

    if i == 1 {
      fmt.Printf("v[%s]\n", v)
    }
  }
}
