package main

import (
	"fmt"
	"gocourse/concurrency/pattern/sample2/logger"
	"os"
	"os/signal"
	"time"
)

type device struct {
	off bool
}

func (d *device) Write(p []byte) (n int, err error) {

	if d.off {

		//simulate disk problem
		time.Sleep(time.Second)
	}

	fmt.Print(string(p))

	return len(p), nil
}

func main() {

	const grs = 10

	var d device

	l := logger.New(&d, grs)

	for i := 0; i < grs; i++ {
		go func(id int) {
			for {
				l.Write(fmt.Sprintf("%d: log data", id))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}

	// We want to control the simulated disk blocking. Capture
	// interrupt signals to toggle device issues. Use <ctrl> z
	// to kill the program.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	for {
		<-sigChan

		d.off = !d.off
	}
}
