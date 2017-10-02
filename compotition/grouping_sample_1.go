package main

import "fmt"

type Speaker interface {
	Speak()
}

type Dog struct {
	Name    string
	IsMamal bool
}

func (d *Dog) Speak() {
	fmt.Println("dog speak :", d.Name)
}

type Cat struct {
	Name    string
	IsMamal bool
}

func (c *Cat) Speak() {
	fmt.Println("cat speak :", c.Name)
}

func main() {

	speaker := []Speaker{

		&Dog{
			Name:    "anjing",
			IsMamal: false,
		},

		&Cat{
			Name:    "kucing",
			IsMamal: true,
		},
	}

	for _, s := range speaker {
		s.Speak()
	}
}

// NOTES:

// Here are some guidelines around declaring types:
// 	* Declare types that represent something new or unique.
// 	* Validate that a value of any type is created or used on its own.
// 	* Embed types to reuse existing behaviors you need to satisfy.
// 	* Question types that are an alias or abstraction for an existing type.
// 	* Question types whose sole purpose is to share common state.
