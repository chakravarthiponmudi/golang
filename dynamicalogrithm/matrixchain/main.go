package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type matrix struct {
	row int
	col int
}

func generateMatrixDimensions(chainSize int) []matrix {
	matrixDimensions := make([]matrix, chainSize)
	generator := rand.New(rand.NewSource(92311))
	seed := chainSize
	matrixDimensions[0].row = generator.Intn(seed) + 1
	for i := range matrixDimensions {

		matrixDimensions[i].col = generator.Intn(seed) + 1
		if i < len(matrixDimensions)-1 {
			matrixDimensions[i+1].row = matrixDimensions[i].col
		}
	}
	return matrixDimensions
}

type dArray = []int

func matrixChainMultiplication(p []matrix, length int) {
	m := make([]dArray, length)
	s := make([]dArray, length)
	for i := range m {
		m[i] = make([]int, length)
		s[i] = make([]int, length)
	}

	for i := 0; i < length; i++ {
		m[i][i] = 0
	}
}

func main() {
	chainLength, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Panic("Error in Chain length", err)
	}

	p := generateMatrixDimensions(chainLength)
	fmt.Println(p)

	matrixChainMultiplication(p, chainLength)
}
