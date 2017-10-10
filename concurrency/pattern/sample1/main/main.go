package main

import (
	"gocourse/concurrency/pattern/sample1/chat"
	"log"
	"os"
	"os/signal"
)

func main() {

	cr := chat.New()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	<-sigChan

	log.Println("Shutting Down Started")
	cr.Close()
	log.Println("Shutting Down Completed")
}
