package main

import (
  "encoding/json"
  "fmt"
)

type user struct {
  Id int
  Name string
}

func main() {

  u, err := retrieveUser("andi")
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Printf("%+v\n", *u)
}

func retrieveUser(name string) (*user, error) {

  r, err := getUser(name)
  if err != nil {
    return nil, err
  }

  var u user
  err = json.Unmarshal([]byte(r), &u)
  return &u, err
}

func getUser(name string) (string, error) {
  response := `{"Id":1432, "Name":"andi"}`
  return response, nil
}
