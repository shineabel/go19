package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg10 sync.WaitGroup
	wg10.Add(10)
	for i := 0; i <5 ;i++{
		fmt.Printf("%d\n",i)
		//fmt.Printf("---------\n")
		go func() {
			fmt.Printf("closure ===%d\n",i)
			wg10.Done()
		}()

		go func(index int) {
			fmt.Printf("===%d\n",index)
			wg10.Done()
		}(i)
	}
	wg10.Wait()
}
