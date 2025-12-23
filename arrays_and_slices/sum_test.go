package main

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 3 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		act := Sum(numbers)
		exp := 1 + 2 + 3

		if act != exp {
			t.Errorf("Actual was %d when expected %d when given %v", act, exp, numbers)
		}
	})

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		act := Sum(numbers)
		exp := 15

		if act != exp {
			t.Errorf("Actual was %d and expected %d when given %v", act, exp, numbers)
		}
	})
}
