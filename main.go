package main

import (
	"fmt"
)

type Sum interface {
	Add(a, b int) int
}
type Sumer struct {
	tt int
}

func (m Sumer) Add(a, b int) int {
	return a + b
}

func main() {
	addr := Sumer{
		tt: 1,
	}
	m := Sum(addr)
	q := m.Add(10, 20) //实现了接口调用(并且runtime.convT32主动在堆上分配内存)
	/*
		test.go:21:13: q escapes to heap
	*/
	fmt.Println(q)
}
