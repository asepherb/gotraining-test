package main

import (
	"fmt"
	"math/rand"
	"time"
)

type data struct {
	UserID string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	//sendRecvUnBuffered()
	// signalClose()
	// closeRange()
	// fanOut()
	// fanOutSem()
	// selectCancel()
	selectDrop()
}

func sendRecvUnBuffered() {

	fmt.Println("** basic sendRecvUnBuffered")

	ch := make(chan data)

	go func() {
		fmt.Println("g2 run")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("g2 initate send")
		ch <- data{"123"}
		fmt.Println("g2 : send ack")
	}()

	fmt.Println("g1 initate recv")
	v := <-ch
	fmt.Println("g1 recieved: ", v)

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// signalClose shows how to close a channel to signal an event.
func signalClose() {

	fmt.Println("** signalClose")

	ch := make(chan struct{})

	go func() {

		fmt.Println("g2 run")
		time.Sleep(100 * time.Millisecond)
		fmt.Println("g2 initiate ")
		close(ch)
		fmt.Println("g2 close ack")
	}()

	fmt.Println("g1 initate recv")
	_, wd := <-ch

	fmt.Println("g1 recieved: ", wd)
	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// closeRange shows how to use range to receive values and
// how using close terminates the loop.
func closeRange() {

	fmt.Println("** closeRange")

	ch := make(chan data)

	const grs = 2

	for g := 0; g < grs; g++ {

		go func(id int) {
			for v := range ch {
				fmt.Println("g2 recieved :", v, " on :", id)
			}
			fmt.Println("g2 : recieved close on :", id)
		}(g)
	}

	const regs = 10
	for i := 0; i < regs; i++ {
		ch <- data{fmt.Sprintf("%d", i)}
	}
	close(ch)

	fmt.Println("g1 close ack")
	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// fanOut shows how to use the fan out pattern to get work
// done concurrently.
func fanOut() {

	fmt.Println("** fanout")

	grs := 20
	ch := make(chan data, grs)

	for g := 0; g < grs; g++ {
		go func(i int) {
			fmt.Println("go running", i)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
			ch <- data{fmt.Sprintf("%d", i)}
			fmt.Println("go end running", i)
		}(g)
	}

	for grs > 0 {

		d := <-ch
		fmt.Println(d)
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// fanOutSemaphore shows how to use the fan out pattern to get work
// done concurrently but limiting the number of active routines.
func fanOutSem() {

	fmt.Println("** fanOutSem")

	grs := 20
	ch := make(chan data, grs)

	const cap = 5
	sem := make(chan bool, cap)

	for g := 0; g < grs; g++ {

		go func(i int) {
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
				ch <- data{fmt.Sprintf("%d", i)}
			}
			<-sem
		}(g)
	}

	for grs > 0 {

		d := <-ch
		fmt.Println("recieved: ", d, len(sem))
		grs--
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// selectCancel shows how to use the select statement to wait for a
// specified amount of time to receive a value.
func selectCancel() {

	fmt.Println("** selectCancel")

	//need to use buffered channel here, if not will be deadlock on g2, for test remove the buffer size
	ch := make(chan data, 1)

	go func() {
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- data{"123"}
		fmt.Println("g2 send ack")
	}()

	tc := time.After(100 * time.Millisecond)

	select {
	case v := <-ch:
		fmt.Println("g1 received : ", v)
	case t := <-tc:
		fmt.Println("g1 timeout", t)
	}
	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}

// selectDrop shows how to use the select to walk away from a channel
// operation if it will immediately block.
func selectDrop() {
	fmt.Println("** selectDrop")

	const cap = 5
	ch := make(chan data, cap)

	go func() {
		for v := range ch {
			fmt.Println("g2 recieved : ", v)
		}
	}()

	const reqs = 20
	for i := 0; i < reqs; i++ {

		select {
		case ch <- data{fmt.Sprintf("%d", i)}:
			fmt.Println("g1 send ack")
		default:
			fmt.Println("g1 drop")
		}
	}

	close(ch)
	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------")
}
