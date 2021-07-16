package main

import (
	"fmt"
	"playground/stack"
)

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	s := stack.NewStack()

	for index, value := range temperatures {
		top, err := s.Peek()
		if err != nil || temperatures[top] >= value {
			s.Push(index)
		} else {
			for {
				top, err = s.Peek()
				if err != nil || temperatures[top] >= value {
					s.Push(index)
					break
				}
				s.Pop()
				result[top] = index - top
			}
		}
	}

	return result
}

func main() {
	var temperatures, expected, result []int

	temperatures = []int{73, 74, 75, 71, 69, 72, 76, 73}
	expected = []int{1, 1, 4, 2, 1, 1, 0, 0}
	fmt.Println(temperatures)
	fmt.Println("expected", expected)
	result = dailyTemperatures(temperatures)
	fmt.Println("result: ", result)

	temperatures = []int{30, 40, 50, 60}
	expected = []int{1, 1, 1, 0}
	fmt.Println(temperatures)
	fmt.Println("expected", expected)
	result = dailyTemperatures(temperatures)
	fmt.Println("result: ", result)

	temperatures = []int{30, 60, 90}
	expected = []int{1, 1, 0}
	fmt.Println(temperatures)
	fmt.Println("expected", expected)
	result = dailyTemperatures(temperatures)
	fmt.Println("result: ", result)
}
