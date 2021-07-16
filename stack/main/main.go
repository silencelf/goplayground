package main

import (
    "fmt"
    "playground/stack"
)

func dailyTemperatures(temperatures []int) []int {
    s := stack.NewStack()
    for !s.IsEmpty() {
        fmt.Println(s.Pop())
    }

    for _, i := range(temperatures) {
        top, err := s.Peek()
        if err != nil || top >= i {
            s.Push(i)
        }
    }
    fmt.Println("stack: ", s)

    return temperatures
}

func main() {
    temperatures := []int{73,74,75,71,69,72,76,73}
    fmt.Println(temperatures)
    expected := []int{1,1,4,2,1,1,0,0}
    fmt.Println(expected)
    result := dailyTemperatures(temperatures)
    fmt.Println(result)
}
