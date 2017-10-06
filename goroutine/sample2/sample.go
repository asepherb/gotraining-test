package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {

	runtime.GOMAXPROCS(2)
}

func main() {

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("start gourutine")

	go func() {
		lowercase()
		wg.Done()
	}()

	go func() {
		upercase()
		wg.Done()
	}()

	fmt.Println("waiting to finish")
	wg.Wait()

	fmt.Println("terminating program")
}

func lowercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'a'; rune < 'a'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}
}

func upercase() {

	// Display the alphabet three times
	for count := 0; count < 3; count++ {
		for rune := 'A'; rune < 'A'+26; rune++ {
			fmt.Printf("%c ", rune)
		}
	}
}
