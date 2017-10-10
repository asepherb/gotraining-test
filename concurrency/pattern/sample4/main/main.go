package main

import (
	"gocourse/concurrency/pattern/sample4/task"
	"log"
	"sync"
	"time"
)

var names = []string{
	"steve",
	"bob",
	"mary",
	"jason",
}

type namePrinter struct {
	name string
}

func (m namePrinter) Work() {
	log.Println(m.name)
	time.Sleep(3 * time.Second)
}

func main() {
	const routine = 10

	t := task.New(routine)

	var wg sync.WaitGroup
	wg.Add(routine * len(names))

	for i := 0; i < routine; i++ {

		for _, name := range names {

			np := namePrinter{
				name: name,
			}

			go func() {
				t.Do(np)
				wg.Done()
			}()
		}
	}

	wg.Wait()

	t.ShutDown()
}
