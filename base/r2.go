package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
	"time"
)

var (
	counter int64
	wg sync.WaitGroup
	shoudown int64
)

func main() {
	wg.Add(2)
	//go inc(1)
	//go inc(2)
	//
	//fmt.Println("wait routine")
	//wg.Wait()
	//fmt.Printf("end", counter)
	go doWork("a")
	go doWork("b")

	time.Sleep(1 * time.Second)
	fmt.Printf("shutdown now")
	atomic.StoreInt64(&shoudown,1)

	wg.Wait()

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

func doWork(name string)  {
	defer wg.Done()
	for {

		fmt.Printf("do work ,name:%s \n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shoudown) == 1 {
			fmt.Printf("shuwdown work: %s \n", name)
			break
		}
	}
}
