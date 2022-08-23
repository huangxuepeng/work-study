package main

import (
	"fmt"
	"time"
)

func main() {
	// c := sync.NewCond(&sync.Mutex{})
	// queue := make([]interface{}, 0, 10)
	// removeFromQueue := func(delay time.Duration) {
	// 	time.Sleep(delay)
	// 	c.L.Lock()
	// 	queue = queue[1:]
	// 	fmt.Println("removed from queue")
	// 	c.L.Unlock()
	// 	c.Signal()
	// }
	// for i := 0; i < 10; i++ {
	// 	c.L.Lock()
	// 	for len(queue) == 2 {
	// 		c.Wait()
	// 	}
	// 	fmt.Println("adding to queue")
	// 	queue = append(queue, struct{}{})
	// 	go removeFromQueue(1 * time.Second)
	// 	c.L.Unlock()
	// }

	/*
		once
	*/
	// var count int
	// in := func() {
	// 	count++
	// }

	// var once sync.Once
	// var ins sync.WaitGroup
	// ins.Add(100)
	// for i := 0; i < 100; i++ {
	// 	go func() {
	// 		defer ins.Done()
	// 		once.Do(in)

	// 	}()
	// }
	// ins.Wait()
	// fmt.Println(count)

	/*
		pool并发模式
	*/
	// myPool := &sync.Pool{
	// 	New: func() interface{} {
	// 		fmt.Println("new instance")
	// 		return struct{}{}
	// 	},
	// }
	// myPool.Get()
	// instance := myPool.Get()
	// fmt.Println(instance)

	// var num int
	// calcPool := &sync.Pool{
	// 	New: func() interface{} {
	// 		num++
	// 		mem := make([]byte, 1024)
	// 		return &mem
	// 	},
	// }
	// calcPool.Put(calcPool.New())
	// calcPool.Put(calcPool.New())
	// calcPool.Put(calcPool.New())
	// calcPool.Put(calcPool.New())
	// numWork := 1024 * 1024 * 1024
	// var wg sync.WaitGroup
	// wg.Add(numWork)
	// for i := numWork; i > 0; i-- {
	// 	go func() {
	// 		defer wg.Done()
	// 		// fmt.Println("1")
	// 		mem := calcPool.Get().(*[]byte)
	// 		defer calcPool.Put(mem)
	// 	}()
	// }
	// wg.Wait()
	// fmt.Println(num)
	// chanOwer := func() <-chan int {

	// 	resultStream := make(chan int, 5)
	// 	go func() {
	// 		defer close(resultStream)
	// 		for i := 0; i <= 10; i++ {
	// 			resultStream <- i
	// 		}
	// 	}()
	// 	return resultStream
	// }
	// res := chanOwer()
	// for result := range res {
	// 	fmt.Println(result)
	// }
	// fmt.Println("结束")
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c)
	}()
	fmt.Println("start...")
	select {
	case <-c:
		fmt.Println("end", time.Since(start))
	}
}
