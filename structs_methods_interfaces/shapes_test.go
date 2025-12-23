package main

import "testing"

func TestPerimeter(t *testing.T) {
	act := Perimeter(10.0, 12.0)
	exp := 44.0

	if act != exp {
		t.Errorf("Actual was %.2f when expected %.2f", act, exp)
	}
}
