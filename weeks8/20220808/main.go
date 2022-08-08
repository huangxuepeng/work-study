package main

import (
	"fmt"
	"sync"
)

func main() {
	// var m sync.Map
	// for i := 0; i < 10; i++ {
	// 	m.Store(i, i+1)
	// }
	// if k, ok := m.Load(2); ok {
	// 	fmt.Println(k.(int))
	// }
	// m.Range(func(k, v interface{}) bool {
	// 	name, age := k.(int), v.(int)
	// 	fmt.Println(name, age)
	// 	return true
	// })
	var m sync.Map
	if a, ok := m.Load(0); ok {
		fmt.Println(a)
	}

}
