package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	actual := Adder(3, 10)
	expected := 13

	fmt.Printf("Actual: %d\nExpected: %d\n", actual, expected)

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
