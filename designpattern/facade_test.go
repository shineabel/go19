package main

import (
	"testing"
	"fmt"
)

func TestNewAPI2(t *testing.T) {
	api := NewAPI2()

	str := api.Test()
	fmt.Println(str)
}
