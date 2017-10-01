package main

import "fmt"

type data struct {
	name string
	age  int
}

func (d data) displayName() {
	fmt.Println("my name is ", d.name)
}

func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "is age", d.age)
}

func main() {

	d := data{
		"andi", 28,
	}

	fmt.Println("proper calls to method")
	d.displayName()
	d.setAge(29)

	fmt.Println("what compiller doing")
	data.displayName(d)
	(*data).setAge(&d, 29)

	fmt.Println("call value reciever with variable")
	f2 := d.setAge
	f2(39)

	d.name = "newAndi"
	f2(39)
}
