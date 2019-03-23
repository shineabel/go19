package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	counter int64
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)
	go inc(1)
	go inc(2)

	fmt.Println("wait routine")
	wg.Wait()
	fmt.Printf("end", counter)

}

func inc( c int)  {

	defer  wg.Done()

	for count := 0 ; count < 2; count++ {
		//value := counter
		//runtime.Gosched()
		//value++
		//counter = value
		atomic.AddInt64(&counter,1)
		runtime.Gosched()
	}
}
