package main

import "fmt"

func main() {
	// Explicit type declaration
	var name string = "Anoop"
	var age int = 25
	var height float64 = 5.9
	var isStudent bool = true

	// Short variable declaration (most common in Go)
	city := "Delhi"
	score := 98.5

	// Multiple assignment
	x, y := 10, 20

	// Constants
	const Pi = 3.14159
	const AppName = "JourneyOfGo"

	fmt.Println("=== Variables Demo ===")
	fmt.Printf("Name: %s, Age: %d\n", name, age)
	fmt.Printf("Height: %.1f, Student: %v\n", height, isStudent)
	fmt.Printf("City: %s, Score: %.1f\n", city, score)
	fmt.Printf("x=%d, y=%d\n", x, y)
	fmt.Printf("Pi=%.5f, App=%s\n", Pi, AppName)

	// Zero values (default values in Go)
	var a int
	var b string
	var c bool
	fmt.Printf("\nZero values: int=%d, string=%q, bool=%v\n", a, b, c)
}
