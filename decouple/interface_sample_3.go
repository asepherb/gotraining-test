package main

import "fmt"

type printer interface {
	print()
}

type user struct {
	name string
}

func (u user) print() {
	fmt.Println("print:", u.name)
}

func main() {

	u := user{"andi"}

	entities := []printer{
		u,
		&u,
	}

	u.name = "apang"

	for _, e := range entities {
		e.print()
	}

	// When we store a value, the interface value has its own
	// copy of the value. Changes to the original value will
	// not be seen.

	// When we store a pointer, the interface value has its own
	// copy of the address. Changes to the original value will
	// be seen.
}
