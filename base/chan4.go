package main

import "fmt"

func calc(routineId int, taskChan chan int, resChan chan int, exitChan chan bool) {
	total := 0
	hint := 0
	for {

		v, isClose := <-taskChan
		if (!isClose) {
			fmt.Printf("task id:%d found task chan closed\n", routineId)
			exitChan <- true
			fmt.Printf("=============task id:%d, total:%d, hint:%d\n",routineId,total,hint)
			return
		} else {
			total++
				flag := true
				for i :=2; i < v; i++ {
					if v%i == 0 {
						flag = false
						break
					}
				}
				if flag {
					resChan <- v
					hint++
					fmt.Printf("number:%d cacl by routine: %d\n",v,routineId)
				}
		}

	}

}

func main() {
	intChan := make(chan int, 1000)
	resultChan := make(chan int, 1000)
	exitChan := make(chan bool, 8)

	go func() {
		for i:=2; i <=1000; i++ {
			intChan <- i
		}

		close(intChan)
	}()


	go func() {
		for i := 0; i < 8; i++ {
			go calc(i,intChan, resultChan, exitChan)
		}
	}()

	go func() {
		for i := 0; i < 8; i++ {
			<- exitChan
		}
		close(resultChan)
	}()


	var count = 0
	for v := range resultChan {
		fmt.Println(v)
		count++
	}

	fmt.Printf("counter %d:",count)
}