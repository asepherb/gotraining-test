package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64

func main() {

	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {

		go func() {

			for count := 0; count < 2; count++ {

				atomic.AddInt64(&counter, 1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("all done", counter)
}
