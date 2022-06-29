package main

import (
	"strconv"
	"testing"
)

func BenchmarkName(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if strconv.Itoa(42) != "42" {
			b.Fatalf("failed")
		}	
	}
}