package main

import "testing"

func BenchmarkEx30(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30)
	}
}
func BenchmarkEx35(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(35)
	}
}
func BenchmarkEx40(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(40)
	}
}
