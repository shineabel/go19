package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutine")

	go func() {
		defer  wg.Done()
		for i := 0; i < 3;i++ {
			for char := 'a'; char < 'a'+26; char ++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c", char)
			}
		}
	}()

	fmt.Println("wait routine finish")
	wg.Wait()

	fmt.Println("end")

}
