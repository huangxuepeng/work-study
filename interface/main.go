package main

import "fmt"

/*
	接口不仅让代码变的更具有通用性和拓展性
*/

type API interface {
	area() float64
	Ll() int
}

type san struct {
	a, b int
	API
}
type ju struct {
	a, b int
	API
}

func (s san) area() float64 {
	return float64(s.a * s.b)
}

func (s san) Ll() int {
	return s.a + s.b
}
func (s san) ll() int {
	return s.a
}
func main() {

	aa := san{
		a: 1,
		b: 2,
	}
	fmt.Println(aa.area())
}
