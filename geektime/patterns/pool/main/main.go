package main

import (
	"github.com/jasonlift/gophercises/geektime/patterns/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	maxGoroutines = 25 // goroutine size
	pooledResource = 2 // resources size
)

type dbConnection struct {
	ID int32
}

func (dbConn * dbConnection) Close() error {
	log.Println("Close: Connection", dbConn.ID)
	return nil
}

var idCounter int32

func createConnection() (io.Closer, error) {
	id := atomic.AddInt32(&idCounter, 1)
	log.Println("Create: New Connection", id)
	return &dbConnection{id}, nil
}

func main() {
	var wg sync.WaitGroup
	wg.Add(maxGoroutines)

	// create a connection pool for management
	apool, err := pool.New(createConnection, pooledResource)
	if err != nil {
		log.Println(err)
	}
	
	// query with each goroutine
	for query := 0; query < maxGoroutines; query++ {
		go func(q int) {
			performQueries(q, apool)
			wg.Done()
		}(query)
	}
	
	wg.Wait()
	log.Println("Shutdown Program.")
	apool.Close()
}

func performQueries(query int, p *pool.Pool) {
	conn, err := p.Acquire()
	if err != nil {
		log.Println(err)
		return
	}

	defer p.Release(conn)

	// mock query and respond
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Microsecond)
	log.Println("QID[%d] CID[%d]\n", query, conn.(*dbConnection).ID)
}
