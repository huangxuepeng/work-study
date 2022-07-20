package main

import (
	"container/heap"
)

type arr []int

func (h arr) Len() int {
	return len(h)
}
func (h arr) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h arr) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 返回最后一个数据
func (h *arr) Pop() interface{} {
	v := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return v
}
func (h *arr) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func main() {
	a := arr{1, 2, 3, 4, 5}
	// 初始化堆
	heap.Init(&a)
	// fmt.Println(a)
	// // heap.Push(&a, 3)
	// // fmt.Println(a)
	// heap.Pop(&a)
	// fmt.Println(a)
	heap.Fix(&a, 1)
}
