package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	ifContinue := true

	for ifContinue {
		fmt.Println("-----SELECT A SHAPE-----")
		fmt.Println("1 - Square\n2 - Rectanglge\n3 - Circle\n0 - Exit")

		fmt.Println("Enter option:")
		scanner.Scan()
		userShape, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		switch userShape {
		case 1:
			GetSquareInfo()
		case 2:
			GetRectInfo()
		case 3:
			GetCircleInfo()
		case 0:
			fmt.Println("Bye now")
			ifContinue = false
		default:
			fmt.Println("error: please select a shape that is listed!")
		}
	}
}

func GetSquareInfo() {

	fmt.Println("Enter the length (whole numbers only):")
	scanner.Scan()
	sqLength, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("1 - Area\n2 - Parameter")
	fmt.Println("Enter what to calculate:")
	scanner.Scan()
	userCalc, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	s := Square{length: sqLength}

	switch userCalc {
	case 1:
		result := s.Area()
		fmt.Println("Area of the Square:", result)
	case 2:
		result := s.Perimeter()
		fmt.Println("Perimeter of the Square:", result)
	default:
		fmt.Println("please select a calculation option that is listed")
	}
}

func GetRectInfo() {
	fmt.Println("Enter the length (whole numbers only):")
	scanner.Scan()
	rectLength, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	fmt.Println("Enter the height(whole numbers only):")
	scanner.Scan()
	rectHeight, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	r := Rectangle{length: rectLength, height: rectHeight} 

	fmt.Println("1 - Area\n2 - Parameter")
	fmt.Println("Enter what to calculate:")
	scanner.Scan()
	userCalc, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	switch userCalc {
	case 1:
		result := r.Area()
		fmt.Println("Area of Rectangle:", result)
	case 2:
		result := r.Perimeter()
		fmt.Println("Perimeter of Rectangle:", result)
	default:
		fmt.Println("ERROR: Please enter a calculation option that is listed.")
	}
}

func GetCircleInfo() {
	fmt.Println("Enter the radius:")
	scanner.Scan()
	cirRadius, err := strconv.ParseFloat(scanner.Text(), 64)
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	c := Circle{radius: cirRadius}

	fmt.Println("1 - Area\n2 - Parameter")
	fmt.Println("Enter what to calculate:")
	scanner.Scan()
	userCalc, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("ERROR:", err)
	}

	switch userCalc {
	case 1:
		result := c.Area()
		fmt.Println("Area of Circle:", result)
	case 2:
		result := c.Perimeter()
		fmt.Println("Perimeter of Circle:", result)
	default:
		fmt.Println("ERROR: please enter a calculation option that is listed")
	}
}

type Shape interface {
	Area() float64
}

// Structs for different shapes
type Square struct {
	length float64
}

type Rectangle struct {
	length float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Perimeter() float64 {
	return 2*(r.length) + 2*(r.height)
}

func (r Rectangle) Area() float64 {
	return r.length * r.height
}

func (s Square) Perimeter() float64 {
	return 4 * s.length
}

func (s Square) Area() float64 {
	return s.length * s.length
}

func (c Circle) Perimeter() float64 {
	return c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

