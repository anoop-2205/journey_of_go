package main

import "fmt"

func main() {
	fmt.Println("=== Arrays & Slices ===")

	// Array — fixed size, rarely used directly
	fmt.Println("\n-- Array --")
	var arr [5]int = [5]int{10, 20, 30, 40, 50}
	fmt.Println("Array:", arr)
	fmt.Printf("Length: %d, Element[2]: %d\n", len(arr), arr[2])

	// Slice — dynamic, built on top of arrays
	fmt.Println("\n-- Slice --")
	fruits := []string{"Apple", "Mango", "Banana", "Cherry"}
	fmt.Println("Slice:", fruits)
	fmt.Printf("Length: %d, Capacity: %d\n", len(fruits), cap(fruits))

	// Slice operations
	fmt.Println("\n-- Slice Operations --")
	fmt.Println("fruits[1:3]:", fruits[1:3]) // Mango, Banana
	fmt.Println("fruits[:2]:", fruits[:2])    // Apple, Mango
	fmt.Println("fruits[2:]:", fruits[2:])    // Banana, Cherry

	// Append to slice
	fruits = append(fruits, "Grape", "Kiwi")
	fmt.Println("After append:", fruits)

	// make() — create slice with length and capacity
	nums := make([]int, 3, 10)
	fmt.Printf("\nmake([]int,3,10): %v, len=%d, cap=%d\n", nums, len(nums), cap(nums))

	// Copy slice
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))
	copied := copy(dst, src)
	fmt.Printf("Copied %d elements: %v\n", copied, dst)

	// 2D slice (matrix)
	fmt.Println("\n-- 2D Slice --")
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, row := range matrix {
		fmt.Println(row)
	}
}
