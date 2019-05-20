package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {

		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {

		Walk(t.Right, ch)
	}

}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		Walk(t1, ch1)
		close(ch1)
	}()

	go func() {
		Walk(t2, ch2)
		close(ch2)
	}()

	for {
		i, more := <-ch1
		j, tmore := <-ch2
		if more && tmore {
			if i != j {
				return false
			}
		} else if more || tmore {
			return false
		} else {
			return true
		}
	}
}

func main() {
	t := tree.New(10)
	// ch := make(chan int, 10)
	// go func() {
	// 	Walk(t, ch)
	// 	close(ch)
	// }()

	// for {
	// 	i, more := <-ch
	// 	if more {
	// 		fmt.Println(i)
	// 	} else {
	// 		return
	// 	}
	// }
	t1 := tree.New(19)

	fmt.Println(t1, t, Same(t1, t))

}
