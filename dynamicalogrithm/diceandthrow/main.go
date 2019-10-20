package main

import "fmt"

type combination struct {
	n  int64
	ts int64
}

var memoizedMap map[combination]int64

var comb []int64
var N int64 = 2500
var K int64 = 6
var sol int64 = 11999

func main() {
	memoizedMap = make(map[combination]int64)
	comb = make([]int64, N)
	count := rolldice(N, K, sol)
	fmt.Printf("%d\n", uint64(count))

}

func clearComb() {
	comb = make([]int64, N)
}

func getfromMemorized(n int64, ts int64) (int64, bool) {
	key := combination{n, ts}
	cnt, ok := memoizedMap[key]
	return cnt, ok
}

func updatememoized(n int64, ts int64, cnt int64) {
	key := combination{n, ts}
	memoizedMap[key] = cnt
}

func rolldice(n int64, k int64, ts int64) int64 {

	var count int64

	cnt, ok := getfromMemorized(n, ts)

	if ok {
		return cnt
	}

	var i int64
	for i = 1; i <= k; i++ {
		if n == N {
			clearComb()
		}
		comb[n-1] = i
		newts := ts - i
		if newts == 0 && n == 1 {
			count = count + 1
			fmt.Println(comb)
		}

		if newts > 0 && n > 1 {
			cnt = rolldice(n-1, k, newts)
			updatememoized(n-1, newts, cnt)
			count = count + cnt
		}

	}
	return count
}
