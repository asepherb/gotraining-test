package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type result struct {
	id  int
	op  string
	err error
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	const routines = 10
	const inserts = routines * 2

	ch := make(chan result, inserts)

	waitInserts := inserts

	for i := 0; i < routines; i++ {

		go func(id int) {

			ch <- insertUser(id)

			ch <- insertTran(id)
		}(i)
	}

	for waitInserts > 0 {

		r := <-ch

		log.Printf("N : %d, ID : %d, OP: %s, Err: %v", waitInserts, r.id, r.op, r.err)

		waitInserts--
	}

	log.Println("insert complete")
}

func insertUser(id int) result {

	r := result{
		id: id,
		op: fmt.Sprintf("insert USERS value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("unable to insert %d into USER table", id)
	}

	return r
}

func insertTran(id int) result {

	r := result{
		id: id,
		op: fmt.Sprintf("insert Trans value (%d)", id),
	}

	if rand.Intn(10) == 0 {
		r.err = fmt.Errorf("unable to insert %d into USER table", id)
	}

	return r
}
