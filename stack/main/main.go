package main

import (
    "fmt"
    "playground/stack"
)

func main() {
    s := stack.NewStack()
    s.Push(1)
    s.Push(2)
    s.Push(3)

    for !s.IsEmpty() {
        fmt.Println(s.Pop())
    }
    fmt.Println(s.Pop())
}
