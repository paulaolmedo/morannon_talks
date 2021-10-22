package main

import (
	"concurrency_and_paralellism/src"
	"fmt"

	"golang.org/x/tour/tree"
)

func main() {
	t1 := tree.New(1)
	t2 := tree.New(2)

	fmt.Printf("t1 equals to t1: %v \n", src.Same(t1, t1))
	fmt.Printf("t1 equals to t2: %v \n", src.Same(t1, t2))
}
