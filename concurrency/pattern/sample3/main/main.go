package main

import (
	"gocourse/concurrency/pattern/sample3/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGourutines = 25
	numPooled     = 2
)

type dbConnection struct {
	ID int32
}

func (dbConn *dbConnection) Close() error {
	log.Println("close connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {

	id := atomic.AddInt32(&idCounter, 1)
	log.Println("create new connection : ", id)

	return &dbConnection{id}, nil
}

func performQuery(query int, p *pool.Pool) {

	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	log.Printf("Query: QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGourutines)

	p, err := pool.New(numPooled, createConnection)
	if err != nil {
		log.Println(err)
		return
	}

	for query := 0; query < maxGourutines; query++ {

		go func(q int) {
			performQuery(q, p)
			wg.Done()
		}(query)
	}

	wg.Wait()
	log.Println("shutdown program")
	p.Close()
}
