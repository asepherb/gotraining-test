package main

import (
	"fmt"
	"reflect"
)

type UnmarshalTypeError struct {
	Value string
	Type  reflect.Type
}

func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + "into value go type" + e.Type.String()
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {

	if e.Type == nil {
		return "json: unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: unmarshal (non pointer " + e.Type.String() + ")"
	}

	return "json: unmarshal (nil " + e.Type.String() + ")"
}

type user struct {
	Name int
}

func main() {

	var u user

	err := Unmarshal([]byte(`{"name":"andi"}`), &u)
	if err != nil {
		switch e := err.(type) {
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError: Value[%s] Type[%v]\n", e.Value, e.Type)

		case *InvalidUnmarshalError:
			fmt.Printf("invalid unmarshal type[%v]\n", e.Type)

		default:
			fmt.Println(err)
		}
		return
	}

	fmt.Println("name : ", u.Name)
}

func Unmarshal(data []byte, v interface{}) error {

	rv := reflect.ValueOf(v)

	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
