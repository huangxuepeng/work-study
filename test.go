package main

import (
	"fmt"
	"sync"
)

type test struct {
	a int
}

func main() {
	var wg sync.WaitGroup
	var l sync.Mutex
	arr := []int{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func(d int) {
			defer wg.Done()
			l.Lock()
			defer l.Unlock()
			arr = append(arr, d)
		}(i)

	}
	wg.Wait()
	fmt.Println(len(arr))
}
