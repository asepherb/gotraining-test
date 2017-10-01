package main

import "fmt"

type reader interface {
	read(b []byte) (int, error)
}

type file struct {
	name string
}

func (file) read(b []byte) (int, error) {
	s := "andi dari file"
	copy(b, s)
	return len(s), nil
}

type pipe struct {
	name string
}

func (pipe) read(b []byte) (int, error) {

	s := "andi dari pipe"
	copy(b, s)
	return len(s), nil
}

func main() {

	f := file{"data.json"}
	p := pipe{"other_service"}

	retrieve(f)
	retrieve(p)

}

func retrieve(r reader) error {

	data := make([]byte, 100)

	len, err := r.read(data)
	if err != nil {
		return err
	}

	fmt.Println(string(data[:len]))
	return nil
}
