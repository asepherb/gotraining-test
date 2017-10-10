package main

import (
	"fmt"
	"sync"
)

var counter int

var mutex sync.Mutex

func main() {

	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {

		go func() {

			for count := 0; count < 2; count++ {

				mutex.Lock()
				{
					value := counter

					value++

					counter = value
				}
				mutex.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("all done", counter)
}
