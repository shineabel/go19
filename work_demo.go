package main

import (
	"github.com/go19/work"
	"log"
	"sync"
	"time"
)

var names = []string{"shine", "codercat", "xianming"}

type namePrinter struct {
	name string
}

func (n *namePrinter) Task() {
	log.Printf(n.name)
	time.Sleep(5 * time.Second)
}
func main() {

	p := work.New(2)

	var wg8 sync.WaitGroup
	wg8.Add(100 * len(names))

	for i := 0; i < 100; i++ {

		for _, n := range names {

			np := namePrinter{
				name: n,
			}

			go func() {
				p.Run(&np)
				wg8.Done()
			}()
		}
	}

	wg8.Wait()
	p.Shudown()

}
