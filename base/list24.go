package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)


var wg4 sync.WaitGroup

var taskCount = 10
var routineCount = 4

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {

	tasks := make(chan  string,10)
	wg4.Add(routineCount)

	for index := 0; index < routineCount; index ++ {
		go work(tasks,index)
	}

	for p := 1; p <= taskCount; p++{
		tasks <- fmt.Sprintf("task:%d",p)
	}
	close(tasks)
	wg4.Wait()

}

func work(tasks chan string, index int)  {

	defer wg4.Done()
	for {
		task, ok := <- tasks
		if !ok {
			fmt.Printf("work %d shutdown\n",index)
			return
		}
		fmt.Printf("work %d start to do work %s\n",index,task)
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("work %d complete task %s\n",index,task)
	}
}