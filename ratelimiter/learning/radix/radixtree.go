package main

import (
	"fmt"

	radix "github.com/armon/go-radix"
)

func main() {
	var tree = radix.New()
	tree.Insert("/", 1)
	tree.Insert("/home", 2)
	tree.Insert("/home/user", 3)
	fmt.Println(tree.Get("/home/use"))
	s, i, b := tree.LongestPrefix("/home/")
	fmt.Println(s, i, b)

}
