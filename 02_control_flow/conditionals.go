package main

import "fmt"

func main() {
	fmt.Println("=== Control Flow: Conditionals ===")

	// Basic if-else
	age := 20
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	// if with init statement (variable scoped to if block)
	if score := 85; score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 75 {
		fmt.Println("Grade: B")
	} else if score >= 60 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Switch statement
	day := "Monday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Printf("%s is a weekday\n", day)
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend\n", day)
	default:
		fmt.Println("Unknown day")
	}

	// Switch with no condition (acts like if-else chain)
	temp := 35
	switch {
	case temp > 40:
		fmt.Println("Very hot!")
	case temp > 30:
		fmt.Println("Hot day")
	case temp > 20:
		fmt.Println("Pleasant")
	default:
		fmt.Println("Cold")
	}
}
