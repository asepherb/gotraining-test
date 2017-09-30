package main

import "fmt"

type user struct {
    id    int
    name  string
}

func main() {

  u1 := user{
    id: 1,
    name: "andi",
  }

  u2 := user{
    id: 2,
    name: "apang",
  }

  display(u1, u2)

  u3 := []user{
    {23, "abe"},
    {32, "pangeran"},
  }

  display(u3...)

  change(u3...)
  fmt.Println("(**************)")
  for _, u := range u3 {
    fmt.Printf("%+v\n", u)
  }
}

func display(users ...user) {
  fmt.Println("(**************)")
  for _, u := range users {
    fmt.Printf("%+v\n", u)
  }
}

func change(users ...user) {
  users[1] = user{99, "change_pangeran"}
}
