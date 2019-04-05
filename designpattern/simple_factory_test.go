package main

import (
	"testing"
	"fmt"
)

func TestType1(t *testing.T)  {

	api := NewAPI(1)
	str := api.say("shine")
	fmt.Println(str)
}

func TestType2(t *testing.T)  {
	api := NewAPI(2)
	str := api.say(" codercat")
	fmt.Println(str)
}