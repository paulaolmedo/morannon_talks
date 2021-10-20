package main

import (
	"fmt"
)

func main() {
	var v int = 19

	fmt.Println("El valor de v antes:", v)

	Increase(*v)

	fmt.Println("El valor de v despues:", v)
}

func Increase(x *int) {
	*x++
	fmt.Println("El valor de v dentro de increase es:", *x)
}
