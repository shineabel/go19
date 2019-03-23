package main

import (
	"sync"
	"fmt"
	"time"
)

var wg3 sync.WaitGroup

func main() {


	c := make(chan int)

	wg3.Add(1)

	go player2(c)

	c <- 1
	wg3.Wait()
}

func player2(c chan  int)  {

	var newRunner int

	runner := <- c
	fmt.Printf("runner %d runing \n",runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("runner %d to the line\n",newRunner )
		go player2(c)
	}

	time.Sleep(5 * time.Second)

	if runner == 4 {
		fmt.Printf(" game end\n")
		wg3.Done()
		return
	}

	fmt.Printf("runner %d exchange to runner %d\n",runner,newRunner)
	c <- newRunner
}