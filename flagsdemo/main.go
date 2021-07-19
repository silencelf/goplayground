package main

import (
    "flag"
    "fmt"
    "log"
)

func main() {
    log.Println("job started...")
    n := flag.Int("n", 1, "the number of runs")
    flag.Parse()
    fmt.Println("number: ", *n)
}
