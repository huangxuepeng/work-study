package main

import (
	"fmt"
	"testing"
)

// func TestMain_Min(t *testing.T) {
// 	uu := min(1, 2)
// 	fmt.Println(uu)
// }

// 压力测试
func BenchmarkHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintf("hello")
	}
}

// 重置定时器
func BenchmarkBigLen(b *testing.B) {
	big := NewBig()

}
