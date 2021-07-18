package main

import "fmt"

type printer interface {
	print()
}

type book struct {
	name string
}

func (b *book) print() {
	fmt.Print(b.name)
}

type game struct {
	name string
}

func (g *game) print() {
	fmt.Println(g.name)
}

func print(p printer) {
	p.print()
}

func main() {
	items := []printer{
		&book{name: "Three Body"},
		&game{name: "Dead Cells"},
	}

	for _, i := range items {
		print(i)
	}
}
