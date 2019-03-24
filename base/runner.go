package main

import (
	"time"
	"fmt"
	"github.com/go19/runner"
	"os"
	"log"
)

const timeout  = 30 * time.Second

func main() {

	fmt.Printf("task start...\n")

	r := runner.New(timeout)

	r.Add(createTask(),createTask(),createTask())

	if err := r.Start(); err != nil {
		switch err {

		case runner.ErrorTimeout:
			fmt.Printf("timeout error")
			os.Exit(1)
		case runner.ErrorInterupt:
			fmt.Printf("interupt error")
			os.Exit(2)
		}
	}
	fmt.Printf("task process end...")
}

func createTask() func(int) {

	return func(id int) {

		log.Printf("processor-task#%d",id)
		time.Sleep(time.Duration(id)*time.Second)
	}
}