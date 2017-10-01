package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u *user) notify() {
	fmt.Println("notify from user embendded", u.name)
}

type admin struct {
	user
	level string
}

func main() {

	ad := admin{
		user: user{
			name: "andi",
		},
		level: "jago",
	}

	ad.notify()
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}
