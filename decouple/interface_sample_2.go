package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name string
}

func (u *user) notify() {
	fmt.Println("sending user to ", u.name)
}

func main() {

	u := user{"andi"}

	//user does not implement notifier (notify method has pointer receiver)
	//sendNotification(u)

	sendNotification(&u)
}

func sendNotification(n notifier) {
	n.notify()
}
