package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Printf("sending user email to %s<%s>\n", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {

	andi := user{"andi", "andi@gmail.com"}
	andi.notify()

	apang := &user{"apang", "apang@gmail.com"}
	apang.notify()

	andi.changeEmail("andinew@gmail.com")
	andi.notify()

	apang.changeEmail("apangnew@gmail.com")
	apang.notify()

	users := []user{
		{"andi", "andi@gmail.com"},
		{"apang", "apang@gmail.com"},
	}

	for _, u := range users {
		u.notify()
	}

}
