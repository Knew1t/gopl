package main

import (
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Echo()
	}
}

func BenchmarkIneffectiveEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IneffectiveEcho()
	}
}
