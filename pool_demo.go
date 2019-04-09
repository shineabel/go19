package main

import (
	"fmt"
	"github.com/go19/pool"
	"io"
	"log"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

var (
	nRoutine       = 25
	nResource uint = 2
	idCounter int32
)

type dbConnection struct {
	id int32
}

func (c *dbConnection) Close() error {
	log.Printf("connection close:id %d", c.id)
	return nil

}

func CreateConnection() (io.Closer, error) {

	i := atomic.AddInt32(&idCounter, 1)
	log.Printf("create new connection %d:", i)
	return &dbConnection{
		id: i,
	}, nil
}

func main() {

	var wg6 sync.WaitGroup

	wg6.Add(nRoutine)
	p, err := pool.New(CreateConnection, nResource)
	if err != nil {
		fmt.Printf("error:", err)
	}

	for q := 0; q < nRoutine; q++ {
		go func(q int) {
			performQuery(q, p)
			wg6.Done()
		}(q)
	}
	wg6.Wait()
	log.Printf("shutdown...")
	p.Close()

}

func performQuery(q int, p *pool.Pool) {

	conn, err := p.Acquire()
	if err != nil {
		fmt.Printf("acquire connection error:", err)
		return
	}

	defer p.Release(conn)

	time.Sleep(time.Duration(rand.Int31n(100)) * time.Millisecond)
	log.Printf("qid %d use  cid %d end ...", q, conn.(*dbConnection).id)
}
