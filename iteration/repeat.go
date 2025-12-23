package main

import (
	"bufio"
	"fmt"
	"strconv"
	"os"
)

func Repeat(char string, length int) string {
	var repeated string
	for i := 0; i < length; i++ {
		repeated += char
	}

	return repeated
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the character to repeat: ")
	scanner.Scan()
	char := scanner.Text()

	fmt.Println("Enter the amount of time to repeat: ")
	scanner.Scan()
	repeatNum, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Errorf("Invalid number: %v", err)
	}

	result := Repeat(char, repeatNum)
	fmt.Println("Result: ", result)
}
