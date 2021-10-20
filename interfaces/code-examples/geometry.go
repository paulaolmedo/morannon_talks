package main

import "fmt"

type shape interface {
	calculateArea() float64
}

type square struct {
	lado float64
}
type triangulo struct {
	base float64
	height float64
}

func (sq square) calculateArea() float64 {
	return sq.lado * sq.lado
}

func (tr triangulo) calculateArea() float64 {
	return tr.base * tr.height
}

func printArea(s shape) {
	fmt.Println(s.calculateArea())
}

func main() {
	tr := triangulo{base: 5, height: 10}
	sq := square{lado: 10}

	printArea(tr)
	printArea(sq)
}