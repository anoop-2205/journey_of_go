package main

import "fmt"

func main() {
	fmt.Println("=== Control Flow: Loops ===")

	// Go only has 'for' — no while, do-while
	// Traditional for loop
	fmt.Println("\n-- Classic for loop --")
	for i := 1; i <= 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// While-style loop
	fmt.Println("\n-- While-style loop --")
	count := 0
	for count < 3 {
		fmt.Printf("count = %d\n", count)
		count++
	}

	// Infinite loop with break
	fmt.Println("\n-- Loop with break --")
	n := 0
	for {
		if n == 3 {
			break
		}
		fmt.Printf("n = %d\n", n)
		n++
	}

	// Continue — skip iteration
	fmt.Println("\n-- Loop with continue (skip even numbers) --")
	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("odd: %d\n", i)
	}

	// Range — iterate over collections
	fmt.Println("\n-- Range over slice --")
	fruits := []string{"Apple", "Mango", "Banana"}
	for index, value := range fruits {
		fmt.Printf("[%d] = %s\n", index, value)
	}

	// Range — iterate over map
	fmt.Println("\n-- Range over map --")
	capitals := map[string]string{"India": "Delhi", "Japan": "Tokyo", "France": "Paris"}
	for country, capital := range capitals {
		fmt.Printf("%s -> %s\n", country, capital)
	}
}
