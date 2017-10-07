package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

const timeoutSeconds = 3 * time.Second

var (
	//recieved os signal
	sigChan = make(chan os.Signal, 1)

	//recieved timeout event
	timeout = time.After(timeoutSeconds)

	complete = make(chan error)

	shutdown = make(chan struct{})
)

func main() {

	log.Println("starting process")

	// We want to receive all interrupt based signals into channel sigChan
	signal.Notify(sigChan, os.Interrupt)

	log.Println("launching processor")
	go processor(complete)

ControlLoop:
	for {

		select {

		case <-sigChan:

			//interup signal recieved by OS
			log.Println("OS INTERRUPT")

			// Close the channel to signal to the processor
			// it needs to shutdown.
			close(shutdown)

			// Set the channel to nil so we no longer process
			// any more of these events.
			sigChan = nil

		case <-timeout:

			log.Println("timout - killing the program")
			os.Exit(1)

		case err := <-complete:

			log.Printf("task complete : error[%s]", err)
			break ControlLoop
		}
	}

	log.Println("process ended")
}

//note arrow is make channel just can recieved
func processor(complete chan<- error) {

	log.Println("Processor - starting")

	var err error

	//remember defer func willbe executed after function done (in this case processor done)
	defer func() {

		if r := recover(); r != nil {
			log.Println("processor - panic", r)
		}

		complete <- err
	}()

	err = doWork()

	log.Println("processor - complete")
}

func doWork() error {

	log.Println("processor - task1")
	time.Sleep(2 * time.Second)

	if checkShutdown() {
		return errors.New("early shutdown")
	}

	log.Println("processor - task2")
	time.Sleep(1 * time.Second)

	if checkShutdown() {
		return errors.New("early shutdown")
	}

	log.Println("processor - task3")
	time.Sleep(1 * time.Second)

	return nil
}

// checkShutdown checks the shutdown flag to determine
// if we have been asked to interrupt processing.
func checkShutdown() bool {
	select {
	case <-shutdown:
		log.Println("checkshutdown - shutdown early")
		return true
	default:
		// If the shutdown channel was not closed,
		// presume with normal processing.
		return false
	}
}
