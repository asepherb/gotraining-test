package main

import "fmt"

type user struct {
    name      string
    surname   string
}

func main() {

  //initiate map
  users := make(map[string]user)

  //set the map
  users["andi"] = user{"andi", "apang"}
  users["apang"] = user{"apang", "apang"}

  for key, value := range users {

    fmt.Println(key, value)
  }

  fmt.Println()

  for key := range users {
    fmt.Println(key)
  }
}
