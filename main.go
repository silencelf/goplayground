package main

import "fmt"

type Number interface {
	int64 | float64
}

func sumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func sumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func main() {
	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic sums: %v and %v \n",
		sumIntsOrFloats(ints),
		sumIntsOrFloats(floats))

	fmt.Printf("Generic sums with type interface: %v and %v \n",
		sumNumbers(ints),
		sumNumbers(floats))
}
