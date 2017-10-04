package main

import (
	"fmt"
	"log"
)

type customError struct{}

func (c *customError) Error() string {
	return "find the bug"
}

// func fail() ([]byte, error) {
//
// 	return nil, nil
// }
func fail() ([]byte, *customError) {

	return nil, nil
}

func main() {

	var err error

	fmt.Println("testing type value stored in interface", err)

	if _, err = fail(); err != nil {

		fmt.Println("why this error", err)
	}

	log.Println("no error")
}
