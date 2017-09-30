package main

import (
	"fmt"
)

type example struct {
	flag    bool
	counter int16
	pi      float32
}

type sample struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {

	var e1 example
	fmt.Printf("%+v\n", e1)

	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}
	fmt.Printf("%+v\n", e2)

	e3 := example{
		flag: true,
	}
	fmt.Printf("%+v\n", e3)

	e4 := example{
		counter: 5,
	}
	fmt.Printf("%+v\n", e4)

	e5 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag: true,
	}

	fmt.Printf("%+v\n", e5)

	//trying to convert diffrent struct
	var s1 sample
	s1 = sample(e1)
	fmt.Printf("%+v\n", s1)
}
