package main

import "testing"

func BenchmarkStart(b *testing.B) {
	for i := 0; i < b.N; i++ {
		start()
	}
}
