package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {
    for i:=0; i < 10; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            fmt.Println("hello")
        }()
    }

    wg.Wait()
    fmt.Println("main ended.")
}

