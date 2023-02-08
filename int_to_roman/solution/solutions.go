package solution

type solution struct {
	values []node
}

type node struct {
	value int
	text  string
}

func New() *solution {
	s := &solution{}
	s.values = make([]node, 13)

	s.values[0] = node{value: 1000, text: "M"}
	s.values[1] = node{value: 900, text: "CM"}
	s.values[2] = node{value: 500, text: "D"}
	s.values[3] = node{value: 400, text: "CD"}
	s.values[4] = node{value: 100, text: "C"}
	s.values[5] = node{value: 90, text: "XC"}
	s.values[6] = node{value: 50, text: "L"}
	s.values[7] = node{value: 40, text: "XL"}
	s.values[8] = node{value: 10, text: "X"}
	s.values[9] = node{value: 9, text: "IX"}
	s.values[10] = node{value: 5, text: "V"}
	s.values[11] = node{value: 4, text: "IV"}
	s.values[12] = node{value: 1, text: "I"}

	return s
}

func (s *solution) Solve(i int) string {
	roman := ""

	for _, v := range s.values {
		for i >= v.value {
			roman += v.text
			i -= v.value
		}
	}
	return roman
}
