package main

import (
	"testing"
)

func BenchmarkPBV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PBV(Obj)
	}
}

func BenchmarkPBP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PBP(&Obj)
	}
}
