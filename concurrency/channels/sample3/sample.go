package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {

	track := make(chan int)

	wg.Add(1)

	go Runner(track)

	track <- 1

	wg.Wait()
}

func Runner(track chan int) {

	const maxExchanges = 4

	var exchange int

	// Wait to receive the baton.
	baton := <-track

	fmt.Printf("runner %d Running with baton\n", baton)

	// New runner to the line.
	if baton < maxExchanges {
		exchange = baton + 1
		fmt.Printf("runner %d to the line\n", exchange)
		go Runner(track)
	}

	//running around the track
	time.Sleep(100 * time.Millisecond)

	//is the race is over
	if baton == maxExchanges {
		fmt.Printf("runner %d finished, race over \n", baton)
		wg.Done()
		return
	}

	fmt.Printf("runner %d exchange with runner %d,\n", baton, exchange)
	track <- exchange
}
