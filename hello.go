package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
	"math/cmplx"
	"net"
	"os"
	"sync"
	"time"
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

/*Pi value*/
const Pi = 3.14

const (
	/*Big number*/
	Big = 1 << 100
	/*Small number*/
	Small = Big >> 100
)

func needsInt(x int) int {
	return x*10 + 1
}

func needsFloat(x float64) float64 {
	return x * 0.1
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

/*Vertex defination*/
type Vertex struct {
	X int
	Y int
}

func reverseBit(n uint) uint {
	rev := bits.Reverse(n)
	fmt.Println(rev)
	return 0
}

func pointers() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

/*
Connect to an HTTP server
*/
func Connect(url string) {
	c, err := net.Dial("tcp", url)
	defer c.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, "GET / HTTP 1.1\r\n\r\n")
	resp, err := bufio.NewReader(c).ReadString('\n')
	fmt.Println(resp)
}

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	arguments := os.Args
	fmt.Println(arguments)

	root := Sqrt(2)
	fmt.Println(root)

	// fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	// fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	// fmt.Printf("Type: %T Value: %v\n", z, z)
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

	v := 1 + 2i
	fmt.Printf("v is of type: %T\n", v)
	fmt.Printf("Pi of type: %T\n", Pi)
	fmt.Println(Pi)

	fmt.Println(needsInt(Small))
	fmt.Println(needsFloat(Small))
	fmt.Println(needsFloat(Big))

	fmt.Println("pointer demo:")
	pointers()

	reversed := reverseBit(43261596)
	fmt.Printf("Reversed: %d.\n", reversed)

	Connect("baidu.com:80")
	SolveKnapsack()

	sc := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.Inc("name")
	}
	time.Sleep(time.Second)
	fmt.Println(sc.Value("name"))
}

/*
SolveKnapsack problem
*/
func SolveKnapsack() {
	values := []float64{60, 100, 120, 130}
	weights := []float64{10, 20, 30, 30}
	capacities := []float64{60.0, 50.0, 40.0}

	fmt.Println(values, weights, capacities)
	for _, capacity := range capacities {
		result := knapsack(capacity, 0, values, weights)
		fmt.Println(result)
	}
}

func knapsack(capacity float64, i int, values, weights []float64) float64 {
	if i >= len(values) {
		return 0
	}
	if capacity < weights[i] {
		return knapsack(capacity, i+1, values, weights)
	}
	v1 := knapsack(capacity-weights[i], i+1, values, weights) + values[i]
	v2 := knapsack(capacity, i+1, values, weights)

	return math.Max(v1, v2)
}
