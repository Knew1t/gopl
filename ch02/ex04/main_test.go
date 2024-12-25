package main

import (
	"math"
	"testing"
)

func benchmark(b *testing.B, num uint64) {
	for i := 0; i < b.N; i++ {
		PopCount3(num)
	}
}

func Benchmark10(b *testing.B) {
	benchmark(b, uint64(math.Pow(2, 10)))
}
func Benchmark20(b *testing.B) {
	benchmark(b, uint64(math.Pow(2, 20)))
}
func Benchmark30(b *testing.B) {
	benchmark(b, uint64(math.Pow(2, 40)))
}
