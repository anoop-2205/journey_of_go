package main

import (
	"fmt"
	"math"
)

// Basic function
func greet(name string) string {
	return "Hello, " + name + "!"
}

// Multiple return values — a Go superpower
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Named return values
func minMax(nums []int) (min, max int) {
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return // naked return — returns named values
}

// Variadic function (variable number of args)
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// First-class functions — functions as values
func apply(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// Closure — function that captures surrounding variables
func makeCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func main() {
	fmt.Println("=== Functions in Go ===")

	fmt.Println(greet("Anoop"))

	result, err := divide(10, 3)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 3 = %.4f\n", result)
	}

	_, err = divide(5, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	nums := []int{3, 1, 7, 2, 9, 4}
	min, max := minMax(nums)
	fmt.Printf("Min: %d, Max: %d\n", min, max)

	fmt.Printf("Sum(1,2,3): %d\n", sum(1, 2, 3))
	fmt.Printf("Sum(1..5): %d\n", sum(1, 2, 3, 4, 5))

	squares := apply([]int{1, 2, 3, 4, 5}, func(n int) int {
		return n * n
	})
	fmt.Println("Squares:", squares)

	// Anonymous function immediately invoked
	circleArea := func(r float64) float64 {
		return math.Pi * r * r
	}(5.0)
	fmt.Printf("Circle area (r=5): %.2f\n", circleArea)

	counter := makeCounter()
	fmt.Printf("Counter: %d, %d, %d\n", counter(), counter(), counter())
}
