package main

import (
	"fmt"
	"strconv"
)

//bottom up approach

type comb = map[string]int64

var d []comb

func bottomUpRollDice(n int, k int, ts int) comb {

	d = make([]comb, 2)
	d[0] = comb{}
	d[1] = comb{}

	for j := 1; j <= k; j++ {
		key := strconv.Itoa(j)
		newts := int64(ts - j)
		if newts >= 0 {
			d[0][key] = newts
		}

	}

	for i := 1; i < n; i++ {
		index := i % 2
		prevIndex := (i + 1) % 2
		fx(d[prevIndex], k, &d[index])
	}

	// fmt.Println(d)
	return d[(n-1)%2]
}

func fx(c comb, k int, r *comb) {

	// c1 := comb{}
	*r = comb{}
	for key, val := range c {
		for j := 1; j <= k; j++ {
			nkey := key + strconv.Itoa(j)
			newts := val - int64(j)
			if newts >= 0 {
				(*r)[nkey] = newts
			}
		}
	}

	// return c1
}

func printsoln(c comb) {
	for key, val := range c {
		if val == 0 {
			fmt.Println(key)
		}
	}
}
func main() {
	ans := bottomUpRollDice(25, 6, 30)
	printsoln(ans)
}
