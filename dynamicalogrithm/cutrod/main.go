package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func topDownCutrod(rodLength int, price []int, r []int, s []int, c int) int {

	if rodLength == 0 {
		return 0
	}

	if r[rodLength-1] > -1 {
		// fmt.Println(r)
		return r[rodLength-1]
	}

	q := -1
	for i := 1; i <= rodLength; i++ {
		value := price[i-1] + topDownCutrod(rodLength-i, price, r, s, c)
		if i < rodLength {
			value = value - c
		}
		if q < value {
			q = value
			r[rodLength-1] = q
			s[rodLength-1] = i
		}
	}
	return q

}

func generatePrice(rodLength int) []int {
	var price = make([]int, rodLength)

	generator := rand.New(rand.NewSource(92311))
	for i := range price {
		seed := int(i)*20 + 10
		price[i] = generator.Intn(seed)
		fmt.Println("Index :", i, " Price : ", price[i])
	}
	return price
}

func printSizes(s []int, rodLength int, p []int) {

	for ; rodLength > 0; rodLength = rodLength - s[rodLength-1] {
		fmt.Printf(" size : %d, price : %d\n", s[rodLength-1], p[s[rodLength-1]-1])
	}
}
func main() {
	lengthOfRod, err := strconv.Atoi(os.Args[1])

	if err != nil {
		log.Panic("Invalid Argument")
	}
	revenue := make([]int, lengthOfRod)
	var sizes = make([]int, lengthOfRod)
	for i := range revenue {
		revenue[i] = -1
		sizes[i] = -1
	}

	price := generatePrice(lengthOfRod)
	cuttingCost, _ := strconv.Atoi(os.Args[2])
	fmt.Println(topDownCutrod(lengthOfRod, price, revenue, sizes, cuttingCost))
	printSizes(sizes, lengthOfRod, price)
}
