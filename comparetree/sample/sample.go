package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// fmt.Println(t)
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	// fmt.Println("value", t.Value)
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

func walkHelper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	fmt.Println("closing channel")
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for {
		i := <-ch1
		j := <-ch2
		fmt.Println(i, "-", j)
		if i != j {
			return false
		}
	}

	return true
}

func main() {
	t := tree.New(10)
	t1 := tree.New(120)
	ch := make(chan int)
	go walkHelper(t, ch)
	for i := range ch {
		fmt.Println(i)
	}
	Same(t, t1)

}
