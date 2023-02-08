package main

import (
	"fmt"
	"playground/int_to_roman/solution"
)

type testCase struct {
	input     int
	exptected string
}

func main() {
	cases := make([]testCase, 3)
	cases[0] = testCase{input: 1, exptected: "I"}
	cases[1] = testCase{input: 4, exptected: "IV"}
	cases[2] = testCase{input: 11, exptected: "XI"}

	s := solution.New()

	for _, v := range cases {
		input := 11
		roman := s.Solve(v.input)

		fmt.Println(input, "converted to roman:", roman, "expected:", v.exptected)
	}
}
