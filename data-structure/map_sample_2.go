package main

import "fmt"

type user struct {
    name      string
    surname   string
}

func main() {

  //declare the map dirrectly
  users := map[string]user{
    "andi":   {"andi", "andi"},
    "apang":  {"apang", "apang"},
    "abe":    {"abe", "abe"},
  }

  for key, value := range users {
    fmt.Println(key, value)
  }

  delete(users, "apang")

  fmt.Println("***************")

  u, found := users["apang"]
  fmt.Println("apang", u, found)
}
