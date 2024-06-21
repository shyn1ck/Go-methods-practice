package main

import "fmt"

type Pair struct {
	X int
	Y int
}

func (p *Pair) Swap() {
	p.X, p.Y = p.Y, p.X
}

func main() {
	pair := Pair{X: 10, Y: 20}
	fmt.Println("Значения до обмена:", pair)
	pair.Swap()
	fmt.Println("Значения после обмена:", pair)
}
