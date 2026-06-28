package main

import "fmt"

func main() {
	fmt.Println("=== Maps in Go ===")

	// Create a map
	capitals := map[string]string{
		"India":   "New Delhi",
		"Japan":   "Tokyo",
		"Germany": "Berlin",
		"Brazil":  "Brasília",
	}

	fmt.Println("\n-- Map contents --")
	for country, capital := range capitals {
		fmt.Printf("%s -> %s\n", country, capital)
	}

	// Access a value
	fmt.Println("\n-- Access --")
	fmt.Println("Capital of India:", capitals["India"])

	// Check if key exists (comma-ok idiom)
	capital, ok := capitals["France"]
	if ok {
		fmt.Println("France:", capital)
	} else {
		fmt.Println("France not found")
	}

	// Add and update
	fmt.Println("\n-- Add & Update --")
	capitals["France"] = "Paris"
	capitals["India"] = "Delhi" // update
	fmt.Println("After update:", capitals["India"])
	fmt.Println("France added:", capitals["France"])

	// Delete a key
	delete(capitals, "Brazil")
	fmt.Println("\nAfter deleting Brazil:")
	fmt.Println("Brazil:", capitals["Brazil"]) // returns zero value ""

	// Map with slice values
	fmt.Println("\n-- Map of slices --")
	teams := map[string][]string{
		"Backend":  {"Anoop", "Raj", "Priya"},
		"Frontend": {"Sara", "Mike"},
	}
	for team, members := range teams {
		fmt.Printf("%s: %v\n", team, members)
	}

	// make() for maps
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	fmt.Println("\nScores:", scores)
	fmt.Println("Total entries:", len(scores))
}
