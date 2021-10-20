package main

import (
	"fmt"
)

func main() {

	v := 19

	var p1 *int
	var p2 = new(int)
	p3 := &v

	fmt.Printf("p1: %T \n", p1)
	fmt.Printf("p2: %T \n", p2)
	fmt.Printf("p3: %T \n", p3)

}
