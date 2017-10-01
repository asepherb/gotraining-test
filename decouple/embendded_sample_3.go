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

func (ad *admin) notify() {
	fmt.Println("notify from admin", ad.name)
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

	ad.user.notify()
}

func sendNotification(n notifier) {
	n.notify()
}
