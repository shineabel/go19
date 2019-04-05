package main

import (
	"testing"
	"fmt"
)

func compute(f OperatorFactory, a int , b int) int {
	o := f.Create()
	o.SetA(a)
	o.SetB(b)
	return o.Result()

}

func TestOperator(t *testing.T)  {

	var f OperatorFactory
	f = PlusOperatorFactory{}
	fmt.Printf("plus result: %d",compute(f,5,4))

	f = MinusOperatorFactory{}
	fmt.Printf("minus result:%d", compute(f,6,4))
}
