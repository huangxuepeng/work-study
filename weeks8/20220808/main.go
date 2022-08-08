package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	// 使用unsafe
	// var arr = *(*[4]int)(unsafe.Pointer(&a[0]))
	// fmt.Printf("%T\n", arr)   // result:[4]int
	// 使用1.17特性
	arr := (*[3]int)(a)
	fmt.Printf("%T", arr) // result: *[4]int
}
