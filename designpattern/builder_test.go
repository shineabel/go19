package main

import (
	"testing"
	"fmt"
)

func TestBuilder1_GetResult(t *testing.T) {
	builder := &Builder1{}
	d := &Director{
		builder:builder,
	}
	d.Construct()
	fmt.Println(builder.GetResult())
}

func TestBuilder2_GetResult(t *testing.T) {
	builder := &Builder2{}
	d := &Director{
		builder:builder,
	}
	d.Construct()
	fmt.Println(builder.GetResult())
}
