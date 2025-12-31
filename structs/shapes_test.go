package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		// Create rectangle object
		rectangle := Rectangle{length: 4, width: 6}

		// Create actual and expected vars
		act := rectangle.Perimeter()
		exp := 20.0

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})

	t.Run("squares", func(t *testing.T) {
		// Create square object
		square := Square{length: 5}

		// Actual and expected vars
		act := square.Perimeter()
		exp := 20.0

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})

	t.Run("circles", func(t *testing.T) {
		// Create circle object
		circle := Circle{radius: 37}

		// Actual and expected variables
		act := circle.Perimeter()
		exp := 37.0

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{length: 4, width: 6}

		act := rect.Area()
		exp := 24.0

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})

	t.Run("squares", func(t *testing.T) {
		square := Square{length: 5}

		act := square.Area()
		exp := 25.0

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{radius: 4}

		act := circle.Area()
		exp := 50.26548245743669

		if act != exp {
			t.Errorf("Actual was %g when expecting %g", act, exp)
		}
	})
}
