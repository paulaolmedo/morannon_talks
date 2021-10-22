// Provides resolution to the excercise given in https://tour.golang.org/concurrency/7
package src

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	preOrder(t, ch)
	defer close(ch)
}

func preOrder(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	preOrder(t.Left, ch)
	preOrder(t.Right, ch)
	ch <- t.Value
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		equalValue := v1 == v2

		if !equalValue {
			return false
		}

		if !ok1 || !ok2 {
			break
		}

	}

	return true
}
