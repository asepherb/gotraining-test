package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u *user) notify() {
	fmt.Println("notify from user", u.name)
}

type admin struct {
	person user //not embendded
	level  string
}

func main() {

	ad := admin{
		person: user{
			name: "andi",
		},
		level: "jago",
	}

	ad.person.notify()
}
