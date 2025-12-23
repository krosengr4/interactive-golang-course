package main

import (
	"testing"
	"fmt"
)

func TestRepeat(t *testing.T) {

	act := Repeat("a", 5)
	exp := "aaaaa"

	fmt.Printf("Expected: %s\nActual: %s\n", exp, act)

	if act != exp {
		t.Errorf("Expected %q but got %q", exp, act)
	}
}

// Benchmarks measure the performance (speed and memory)
// Use the -b flag when wanting to see the benchmark
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 5)
	}
}
