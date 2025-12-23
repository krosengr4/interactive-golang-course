package main

import (
	"fmt"
	"strings"
)

func Perimeter(width, height float64) float64 {
	return (2 * width) + (2 * height)
}

func main() {
	h := 5.0
	w := 3.0

	fmt.Printf("Height is %.2f\nWidth is %.2f\n", h, w)

	result := Perimeter(h, w)
	fmt.Println(strings.Repeat("-", 30))
	fmt.Printf("The perimeter is: %.2f\n", result)
}
