package main

import "testing"

func BenchmarkDir(b *testing.B) {
	add := Sumer{
		tt: 1,
	}
	for i := 0; i < b.N; i++ {
		add.Add(10, 20)
	}
}

func BenchmarkInterface(b *testing.B) {
	add := Sumer{
		tt: 1,
	}
	for i := 0; i < b.N; i++ {
		Sum(add).Add(10, 20)
	}
}
