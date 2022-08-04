// package main

// import "fmt"

// func main() {
// 	a := 0
// 	ch1, ch2 := make(chan int, 10), make(chan int, 1)
// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			if i > 5 {
// 				ch1 <- i
// 			} else {
// 				ch2 <- i
// 			}
// 		}
// 	}()
// 	for {
// 		a++
// 		select {
// 		case <-ch1:
// 			fmt.Println("结束了哈")
// 			return
// 		case <-ch2:
// 			a++
// 		default:
// 		}
// 		if a == 4 {
// 			fmt.Println("没错了")
// 			return
// 		}
// 	}
// }

package main

import "fmt"

func main() {
	a, b := make(chan int, 1), make(chan int, 4)
	a <- 1
	b <- 2
	test(a, b)
}

func test(a chan int, b chan int) {
	fmt.Println(<-a, <-b)
}
