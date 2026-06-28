package main

import (
	"fmt"
	"math"
)

// Struct definition
type Person struct {
	Name string
	Age  int
	City string
}

// Method on struct (value receiver)
func (p Person) Greet() string {
	return fmt.Sprintf("Hi, I'm %s from %s, age %d", p.Name, p.City, p.Age)
}

// Method with pointer receiver (can modify the struct)
func (p *Person) Birthday() {
	p.Age++
}

// Embedding — Go's composition (not inheritance)
type Employee struct {
	Person      // embedded struct
	CompanyName string
	Salary      float64
}

func (e Employee) Info() string {
	return fmt.Sprintf("%s works at %s earning %.0f/month", e.Name, e.CompanyName, e.Salary)
}

// Interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Function accepting the interface
func printShapeInfo(s Shape) {
	fmt.Printf("Area: %.2f, Perimeter: %.2f\n", s.Area(), s.Perimeter())
}

func main() {
	fmt.Println("=== Structs & Interfaces ===")

	// Create struct
	p1 := Person{Name: "Anoop", Age: 25, City: "Delhi"}
	fmt.Println(p1.Greet())
	p1.Birthday()
	fmt.Printf("After birthday: age = %d\n", p1.Age)

	// Embedding
	fmt.Println("\n-- Embedding --")
	emp := Employee{
		Person:      Person{Name: "Raj", Age: 30, City: "Mumbai"},
		CompanyName: "TechCorp",
		Salary:      75000,
	}
	fmt.Println(emp.Info())
	fmt.Println(emp.Greet()) // promoted method from Person

	// Interfaces
	fmt.Println("\n-- Interfaces --")
	c := Circle{Radius: 5}
	r := Rectangle{Width: 4, Height: 6}

	fmt.Print("Circle: ")
	printShapeInfo(c)
	fmt.Print("Rectangle: ")
	printShapeInfo(r)

	// Slice of interface — polymorphism
	shapes := []Shape{Circle{3}, Rectangle{2, 8}, Circle{7}}
	fmt.Println("\nAll shapes:")
	for i, s := range shapes {
		fmt.Printf("  [%d] Area=%.2f\n", i, s.Area())
	}
}
