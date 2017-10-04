package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type AppError struct {
	State int
}

func (a *AppError) Error() string {

	return fmt.Sprintf("app error, state : %d", a.State)
}

func main() {

	if err := firstCall(10); err != nil {

		switch v := errors.Cause(err).(type) {
		case *AppError:
			fmt.Println("custom app error:", v.State)
		default:
			fmt.Println("default error")
		}

		fmt.Println("\nStack Trace\n*************")
		fmt.Printf("%+v", err)
	}
}

func firstCall(i int) error {

	if err := secondCall(i); err != nil {
		return errors.Wrapf(err, "secondCall")
	}
	return nil
}

func secondCall(i int) error {

	if err := thirdCall(); err != nil {
		return errors.Wrapf(err, "thirdcall")
	}
	return nil
}

func thirdCall() error {
	return &AppError{99}
}
