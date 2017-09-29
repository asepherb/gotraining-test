package main

type user struct {
    name string
    email string
    logins int
}

func main(){

  stayOnStack()
  escapeToHeap()
}

func stayOnStack() user {

  u := user{
    name: "Bill",
    email: "bill@gmail.com",
  }

  return u
}

func escapeToHeap() *user {

  u := user{
    name: "Bill",
    email: "bill@gmail.com",
  }

  return &u
}
