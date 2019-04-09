package main

import "fmt"

func main() {

	total := 100000
	count := 0

	for {
		if total > 0 {
			if total <= 50000 {
				total = total - 1000
			} else {
				total = total - 100000*0.05
			}
			count++
			fmt.Println("current total", total, count)
		} else {
			break
		}
	}

	fmt.Println("count :", count)
}
