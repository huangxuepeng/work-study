package main

import (
	"sync"
)

type test struct {
	a int
}

func main() {
	var wg sync.WaitGroup
	b := &test{
		a: 10,
	}

	for i := 0; i < 1000; i++ {
		wg.Add(2)

		go func(d int) {
			defer wg.Done()
			b.a = d
		}(i)
		go func() {
			defer wg.Done()
			b.a = 0
		}()
	}
	wg.Wait()
}
