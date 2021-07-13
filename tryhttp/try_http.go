package main

import (
    "fmt"
//    "net/http"
//    "io"
//    "html"
//    "log"
)

func main() {
    num, sum := 30, 1.0

    for i:=0; i< num; i++ {
        sum *= (365 - float64(i))/365
    }

    fmt.Println(num, sum)
}
