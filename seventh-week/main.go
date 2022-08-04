package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	// 实现异步操作
	// sum := <-test(2, 3)
	// fmt.Println(sum)
	test(2, 3, func(sum int) {
		fmt.Println(sum)
	})
}

func test(a, b int, call func(sum int)) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		sum := a + b
		fmt.Println("1")
		call(sum)
	}()
	wg.Wait()
}
