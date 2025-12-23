package main

import "fmt"

func main() {
	numbers := []int{2, 2, 3, 3, 4}

	result := Sum(numbers)
	fmt.Println("Result:", result)

}

func Sum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}
