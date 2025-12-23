package integers

import "testing"

func TestAdder(t *testing.T) {
	actual := Adder(3, 10)
	expected := 13

	if actual != expected {
		t.Errorf("Expected %d but got %d", expected, actual)
	}
}
