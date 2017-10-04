package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct{}

func (car) String() string {
	return "VROOM"
}

type cloud struct{}

func (cloud) String() string {
	return "BIG DATA"
}

func main() {

	rand.Seed(time.Now().UnixNano())

	msv := []fmt.Stringer{
		car{},
		cloud{},
	}

	for i := 0; i < 10; i++ {

		rn := rand.Intn(2)

		if _, is := msv[rn].(cloud); is {
			fmt.Println("Luck")
			continue
		}

		fmt.Println("Unluck")
	}
}
