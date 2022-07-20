package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	// a := l.Len()
	// fmt.Println(a)
	for i := 0; i < 10; i++ {
		l.PushBack(i)
	}
	a := l.Len()
	fmt.Println(a)
}
