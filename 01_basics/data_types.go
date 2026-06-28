package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=== Data Types in Go ===")

	// Integer types
	var i8 int8 = 127          // -128 to 127
	var i16 int16 = 32767      // -32768 to 32767
	var i32 int32 = 2147483647 // ~2 billion
	var i64 int64 = 9223372036854775807
	var u uint = 42 // unsigned (no negative)

	fmt.Println("\n-- Integers --")
	fmt.Printf("int8: %d, int16: %d, int32: %d\n", i8, i16, i32)
	fmt.Printf("int64: %d, uint: %d\n", i64, u)

	// Float types
	var f32 float32 = 3.14
	var f64 float64 = math.Pi

	fmt.Println("\n-- Floats --")
	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %.15f\n", f64)

	// String & rune (character)
	var greeting string = "Namaste Go!"
	var ch rune = 'G' // rune is alias for int32

	fmt.Println("\n-- String & Rune --")
	fmt.Printf("string: %s\n", greeting)
	fmt.Printf("rune: %c (value: %d)\n", ch, ch)
	fmt.Printf("string length: %d bytes\n", len(greeting))

	// Type conversion (explicit in Go — no implicit casting)
	var myInt int = 42
	var myFloat float64 = float64(myInt)
	var backToInt int = int(myFloat)
	fmt.Printf("\nConversion: int(%d) -> float64(%f) -> int(%d)\n", myInt, myFloat, backToInt)
}
