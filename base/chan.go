package main

import (
	"math/rand"
	"time"
	"sync"
	"fmt"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg2 sync.WaitGroup

func main() {

	court := make(chan  int)

	wg2.Add(2)

	go player("shine",court)
	go player("cat",court)

	fmt.Println("game start")

	court <- 1

	wg2.Wait()

}

func player(name string, court chan int)  {

	defer wg2.Done()

	for {
		ball, ok := <- court
		if !ok {
			fmt.Printf(" %s win \n ",name)
			return
		}

		n := rand.Intn(100)
		if n % 13 == 0{
			fmt.Printf("%s miss \n",name)
			close(court)
			return
		}
		fmt.Printf("%s hit %d \n",name,court)
		ball ++
		court <- ball
	}

}
