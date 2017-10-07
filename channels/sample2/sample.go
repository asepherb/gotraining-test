package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	court := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("andi", court)
		wg.Done()
	}()

	go func() {
		player("apang", court)
		wg.Done()
	}()

	court <- 1

	wg.Wait()
}

func player(name string, court chan int) {
	for {
		// Wait for the ball to be hit back to us.
		ball, wd := <-court

		if !wd {
			fmt.Printf("player %s won \n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("player %s missed \n", name)
			close(court)
			return
		}

		fmt.Printf("player %s hit %d \n", name, ball)
		ball++

		court <- ball
	}
}
