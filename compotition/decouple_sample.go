package main

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Data struct {
	Line string
}

type Puller interface {
	Pull(d *Data) error
}

type Storer interface {
	Store(d *Data) error
}

type System struct {
	Puller
	Storer
}

//pull data
func pull(p Puller, data []Data) (int, error) {

	for i := range data {
		if err := p.Pull(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

//store data
func store(s Storer, data []Data) (int, error) {

	for i := range data {
		if err := s.Store(&data[i]); err != nil {
			return i, err
		}
	}

	return len(data), nil
}

//copy
func Copy(sys *System, batch int) error {

	data := make([]Data, batch)

	for {
		i, err := pull(sys, data)
		if i > 0 {

			if _, err := store(sys, data[:i]); err != nil {
				return err
			}
		}

		if err != nil {

			return err
		}
	}
}

//==========================================================================
type Xenia struct {
	Host    string
	Timeout time.Duration
}

func (*Xenia) Pull(d *Data) error {

	switch rand.Intn(10) {
	case 1, 9:
		return io.EOF

	case 5:
		return errors.New("method not found")

	default:
		d.Line = "Data"
		fmt.Println("in: ", d.Line)
		return nil
	}
}

type Pilar struct {
	Host    string
	Timeout time.Duration
}

func (*Pilar) Store(d *Data) error {
	fmt.Println("out : ", d.Line)
	return nil
}

func main() {

	system := System{

		Puller: &Xenia{},
		Storer: &Pilar{},
	}

	if err := Copy(&system, 3); err != io.EOF {
		fmt.Println(err)
	}
}
